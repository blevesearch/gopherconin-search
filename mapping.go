package main

import (
	"github.com/blevesearch/bleve"
)

func buildMapping() *bleve.IndexMapping {
	enFieldMapping := bleve.NewTextFieldMapping()
	enFieldMapping.Analyzer = "en"

	eventMapping := bleve.NewDocumentMapping()
	eventMapping.AddFieldMappingsAt("summary", enFieldMapping)
	eventMapping.AddFieldMappingsAt("description", enFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = eventMapping
	mapping.DefaultAnalyzer = "en"

	return mapping
}
