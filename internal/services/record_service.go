package services

import (
	"ehr-headless-api/internal/models"
	"ehr-headless-api/internal/repositories"
	"ehr-headless-api/pkg/search"
)

type RecordService interface {
	CreateRecordEntry(entry *models.RecordEntry) error
}

type recordService struct {
	repo   repositories.RecordRepository
	search search.SearchService
}

func NewRecordService(repo repositories.RecordRepository, search search.SearchService) RecordService {
	return &recordService{repo, search}
}

func (s *recordService) CreateRecordEntry(entry *models.RecordEntry) error {
	err := s.repo.CreateRecordEntry(entry)
	if err != nil {
		return err
	}

	go s.search.IndexRecordEntry(entry)

	return nil
}
