package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blevesearch/bleve"
	bleveHttp "github.com/blevesearch/bleve/http"
)

var indexPath = flag.String("index", "gopherconin.bleve", "index path")
var eventsPath = flag.String("events", "gophercon-schedule.html", "fosdem events ical path")
var bindAddr = flag.String("addr", ":8099", "http listen address")
var update = flag.Duration("update", 0, "update every")

var lastUpdated time.Time

func main() {

	flag.Parse()

	// turn on http request logging
	bleveHttp.SetLog(log.New(os.Stderr, "bleve.http", log.LstdFlags))

	// open index
	index, err := bleve.Open(*indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		// or create new if it doesn't exist
		mapping := buildMapping()
		index, err = bleve.New(*indexPath, mapping)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// insert/update index in background
	go batchIndexEvents(index, *eventsPath)

	// update data periodically
	if *update > 0 {
		log.Printf("Updating every: %s", *update)
		ticker := time.NewTicker(*update)
		go func() {
			for _ = range ticker.C {
				log.Printf("Updating now")
				go batchIndexEvents(index, *eventsPath)
			}
		}()
	}

	// start server
	startServer(index, *bindAddr)
}

func startServer(index bleve.Index, addr string) {
	// create a router to serve static files
	router := staticFileRouter()

	// add the API
	bleveHttp.RegisterIndexName("gopherconin", index)
	searchHandler := bleveHttp.NewSearchHandler("gopherconin")
	router.Handle("/api/search", searchHandler).Methods("POST")
	lastHandler := new(lastUpdatedHandler)
	router.Handle("/api/lastUpdated", lastHandler).Methods("GET")

	http.Handle("/", router)
	log.Printf("Listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func batchIndexEvents(index bleve.Index, path string) {
	count := 0
	batch := bleve.NewBatch()
	for event := range parseEvents(path) {
		batch.Index(event.UID, event)
		if batch.Size() >= 100 {
			err := index.Batch(batch)
			if err != nil {
				log.Println(err)
				return
			}
			count += batch.Size()
			log.Printf("Indexed %d Events\n", count)
			batch = bleve.NewBatch()
		}
	}
	if batch.Size() > 0 {
		err := index.Batch(batch)
		if err != nil {
			log.Println(err)
			return
		}
		count += batch.Size()
	}
	log.Printf("Indexed %d Events\n", count)
}
