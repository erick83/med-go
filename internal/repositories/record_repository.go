package repositories

import (
	"ehr-headless-api/internal/database"
	"ehr-headless-api/internal/models"
)

type RecordRepository interface {
	CreateRecordEntry(entry *models.RecordEntry) error
}

type recordRepository struct{}

func NewRecordRepository() RecordRepository {
	return &recordRepository{}
}

func (r *recordRepository) CreateRecordEntry(entry *models.RecordEntry) error {
	result := database.DB.Create(entry)
	return result.Error
}
