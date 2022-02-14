package handlers

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Controller struct {
	DB *sql.DB
}

type Widget struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Controller) Welcome(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "Hello world",
		"code":    "200",
	})
}

func (c *Controller) Widgets(ctx *fiber.Ctx) error {
	rows, err := c.DB.QueryContext(
		context.Background(),
		"select id, name, weight, created_at from widgets limit $1",
		10,
	)
	//
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	defer rows.Close()
	var widgets []Widget
	for rows.Next() {
		widget := Widget{}
		if err := rows.Scan(&widget.Id, &widget.Name, &widget.Weight, &widget.CreatedAt); err != nil {
			return err
		}
		widgets = append(widgets, widget)
	}
	
	return ctx.JSON(widgets)
}
