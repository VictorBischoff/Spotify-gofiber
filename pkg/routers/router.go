package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/handlers"
)

func Routers(app *fiber.App) {
	app.Get("/", handlers.IndexHandler)
}
