package routers

import "github.com/gofiber/fiber/v2"

func AllRouters(app *fiber.App){
	SongRouters(app)
	SpotifyRouters(app)
}