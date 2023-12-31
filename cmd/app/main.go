package main

import (
	"fmt"
	"log"

	"github.com/aonufrei/learn-de-be/internal/data"
	"github.com/aonufrei/learn-de-be/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	log.Println("Creating repository")
	repo := data.CreateRepository("./words.xlsx")
	log.Println("Existing Topics: ")
	for _, t := range repo.Topics {
		fmt.Printf("%d : %s \n", t.ID, t.Name)
	}
	log.Println("Existing Words: ")
	for _, w := range repo.Words {
		fmt.Printf("%d : %d : %d : %s : %s \n", w.ID, w.TopicId, w.Article, w.Word, w.Translation)
	}

	log.Println("Running the application")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	log.Println("Adding Routes")
	handlers.AddRoutes(app, repo)
	app.Listen(":3000")
}
