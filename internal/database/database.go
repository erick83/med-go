package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"ehr-headless-api/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected")
	DB.Logger = DB.Logger.LogMode(logger.Info)
	log.Println("running migrations")
	err = DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.PatientProfile{},
		&models.DoctorProfile{},
		&models.NurseProfile{},
		&models.Clinic{},
		&models.Specialty{},
		&models.Appointment{},
		&models.MedicalRecord{},
		&models.RecordEntry{},
		&models.LabTestOrder{},
		&models.LabTestResult{},
		&models.Prescription{},
		&models.PrescribedMedication{},
	)

	if err != nil {
		log.Fatal("Migration Failed. \n", err)
	}

}
