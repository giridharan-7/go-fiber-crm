package main

import (
	"fmt"

	"github.com/giridharan-7/go-fiber-crm/database"
	"github.com/giridharan-7/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opend to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	/* In this Go code, the app.Listen(3000) function is used to start
	the HTTP server and listen for incoming requests on port 3000. This
	function is provided by the Fiber web framework, which is a popular lightweight
	framework for building web applications in Go. THIS Listen IS FROM FIBER FRAMEWORK
	LISTENANDSERVE IS FROM HTTP PACKAGE */
	defer database.DBConn.Close()
}
