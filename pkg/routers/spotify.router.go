package routers

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/handlers"
)

func SpotifyRouters(app *fiber.App) {
	app.Get("/", adaptor.HTTPHandler(handler(handlers.RegHand)))
	app.Get("/callback", adaptor.HTTPHandler(handler(handlers.CompleteAuth)))
}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}
