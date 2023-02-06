package routers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/handlers"
)

func Routers(app *fiber.App) {
	//index router
	app.Get("/", handlers.IndexHandler)

	app.Get("/api/v1/song", handlers.GetSongs)
	app.Get("/api/v1/song/:id", handlers.GetSong)
	app.Post("/api/v1/song", handlers.AddSong)
	app.Put("/api/v1/song/:id", handlers.UpdateSong)
	app.Delete("/api/v1/song/:id", handlers.DeleteSong)
}
