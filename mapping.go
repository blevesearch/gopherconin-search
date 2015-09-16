package main

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/language/en"
)

func buildMapping() *bleve.IndexMapping {
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
