package search

import (
	"log"

	"ehr-headless-api/internal/models"
)

type SearchService interface {
	IndexRecordEntry(entry *models.RecordEntry)
}

type searchService struct{}

func NewSearchService() SearchService {
	return &searchService{}
}

func (s *searchService) IndexRecordEntry(entry *models.RecordEntry) {
	// In a real implementation, you would use the Elasticsearch client to index the document.
	log.Printf("Indexing RecordEntry ID %d for Patient %d in Elasticsearch", entry.ID, entry.MedicalRecordID)
	// Simulate some work
	// time.Sleep(1 * time.Second)
	log.Printf("Successfully indexed RecordEntry ID %d", entry.ID)
}
