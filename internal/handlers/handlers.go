package handlers

import (
	"github.com/aonufrei/learn-de-be/internal/data"
	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App, repo *data.Repository) {
	app.Get("/topics", GetAllTopicsHandler(repo))
	app.Get("/topic/:topicId/words", GetWordsByTopicHandler(repo))
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func GetAllTopicsHandler(repo *data.Repository) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(repo.GetAllTopics())
	}
}

func GetWordsByTopicHandler(repo *data.Repository) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		topicId, err := c.ParamsInt("topicId")
		if err != nil {
			c.JSON(errorResponse{400, "Topic id is required"})
		}
		return c.JSON(repo.GetWordsByTopic(topicId))
	}
}
