package main

import (
	"fmt"
	_ "github.com/Hsun-Weng/human-resource-service/docs"
	"github.com/Hsun-Weng/human-resource-service/internal/config"
	"github.com/Hsun-Weng/human-resource-service/internal/constants/job_role"
	models2 "github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/internal/routers"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// @title Human resource service
// @version 1.0
// @description The simple hr management service

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {
	db := util.NewDb()

	setupDatabase(db)

	router, _ := routers.InitializeEngine()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.GetEnv(config.Port, "8080")),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func setupDatabase(db *gorm.DB) {
	migrate(db)
	seed(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models2.Employee{}, &models2.Leave{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	fmt.Println("Database connected and migrated successfully")
}

func seed(db *gorm.DB) {
	var count int64
	db.Model(&models2.Employee{}).Count(&count)

	generalPassword, _ := util.HashPassword("123456")
	if count == 0 {
		employees := []models2.Employee{
			{
				CompanyEmail:  "john.doe@company.com",
				Name:          "John Doe",
				PersonalEmail: "john.doe@gmail.com",
				Password:      generalPassword,
				Phone:         "123-456-7890",
				Role:          string(job_role.Employee),
				LivingAddress: "123 Main St, Springfield, IL",
				Salary:        30000,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				CompanyEmail:  "jane.smith@company.com",
				Name:          "Jane Smith",
				PersonalEmail: "jane.smith@gmail.com",
				Password:      generalPassword,
				Phone:         "234-567-8901",
				Role:          string(job_role.Employee),
				LivingAddress: "456 Oak St, Springfield, IL",
				Salary:        30000,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				CompanyEmail:  "bob.johnson@company.com",
				Name:          "Bob Johnson",
				PersonalEmail: "bob.johnson@gmail.com",
				Password:      generalPassword,
				Phone:         "345-678-9012",
				Role:          string(job_role.Employee),
				LivingAddress: "789 Pine St, Springfield, IL",
				Salary:        30000,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				CompanyEmail:  "alice.williams@company.com",
				Name:          "Alice Williams",
				PersonalEmail: "alice.williams@gmail.com",
				Password:      generalPassword,
				Role:          string(job_role.Manager),
				Phone:         "456-789-0123",
				LivingAddress: "101 Maple St, Springfield, IL",
				Salary:        50000,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
			{
				CompanyEmail:  "charlie.brown@company.com",
				Name:          "Charlie Brown",
				PersonalEmail: "charlie.brown@gmail.com",
				Password:      generalPassword,
				Role:          string(job_role.Manager),
				Phone:         "567-890-1234",
				LivingAddress: "202 Birch St, Springfield, IL",
				Salary:        50000,
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			},
		}

		if err := db.Create(&employees).Error; err != nil {
			fmt.Println("Error seeding data:", err)
		} else {
			fmt.Println("Seed data inserted successfully.")
		}
	} else {
		fmt.Println("Data already exists. Skipping seed.")
	}
}
