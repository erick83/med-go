package handlers

import (
	"ehr-headless-api/internal/models"
	"ehr-headless-api/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

type RecordHandler struct {
	service services.RecordService
}

func NewRecordHandler(service services.RecordService) *RecordHandler {
	return &RecordHandler{service}
}

type CreateRecordEntryDTO struct {
	MedicalRecordID uint   `json:"medicalRecordId" validate:"required"`
	AuthorID        uint   `json:"authorId" validate:"required"`
	Content         string `json:"content" validate:"required"`
	StructuredData  string `json:"structuredData"` // Assuming JSON string input
}

func (h *RecordHandler) CreateRecordEntry(c *fiber.Ctx) error {
	dto := new(CreateRecordEntryDTO)
	if err := c.BodyParser(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Basic validation placeholder
	// In a real app, use a validator library like go-playground/validator
	if dto.Content == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Content is required"})
	}

	entry := &models.RecordEntry{
		MedicalRecordID: dto.MedicalRecordID,
		AuthorID:        dto.AuthorID,
		Content:         dto.Content,
		StructuredData:  datatypes.JSON(dto.StructuredData),
	}

	if err := h.service.CreateRecordEntry(entry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create record entry"})
	}

	return c.Status(fiber.StatusCreated).JSON(entry)
}
