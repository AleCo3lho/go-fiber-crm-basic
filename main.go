package main

import (
	"fmt"
	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("ERROR: feiled to connect database.")
	}
	fmt.Println("INFO: Connection stablished to database.")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("INFO: Database Migrated.")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(":3000")
}
