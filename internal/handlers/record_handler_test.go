package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"ehr-headless-api/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// MockRecordService is a mock implementation of the RecordService interface
type MockRecordService struct {
	CreateRecordEntryFunc func(entry *models.RecordEntry) error
}

func (m *MockRecordService) CreateRecordEntry(entry *models.RecordEntry) error {
	if m.CreateRecordEntryFunc != nil {
		return m.CreateRecordEntryFunc(entry)
	}
	return nil
}

func TestCreateRecordEntry(t *testing.T) {
	app := fiber.New()

	mockService := &MockRecordService{}
	handler := NewRecordHandler(mockService)

	app.Post("/api/v1/patients/:patientId/history", handler.CreateRecordEntry)

	dto := CreateRecordEntryDTO{
		MedicalRecordID: 1,
		AuthorID:        1,
		Content:         "Test content",
		StructuredData:  `{"key":"value"}`,
	}
	body, _ := json.Marshal(dto)

	req := httptest.NewRequest("POST", "/api/v1/patients/1/history", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)
}
