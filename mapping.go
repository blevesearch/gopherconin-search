package main

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/language/en"
	"github.com/blevesearch/bleve/mapping"
)

func buildMapping() mapping.IndexMapping {
	enFieldMapping := bleve.NewTextFieldMapping()
	enFieldMapping.Analyzer = en.AnalyzerName

	eventMapping := bleve.NewDocumentMapping()
	eventMapping.AddFieldMappingsAt("summary", enFieldMapping)
	eventMapping.AddFieldMappingsAt("description", enFieldMapping)

	mapping := bleve.NewIndexMapping()
	mapping.DefaultMapping = eventMapping
	mapping.DefaultAnalyzer = en.AnalyzerName

	return mapping
}
