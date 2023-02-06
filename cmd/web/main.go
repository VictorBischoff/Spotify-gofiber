package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/victorbischoff/GOFIBER-TMPL/database"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/routers"
)

func main() {

	engine := html.New("./public", ".gohtml")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public", fiber.Static{
		Compress: true,
	})

	database.InitDatabase()

	routers.Routers(app)

	log.Fatal(app.Listen(":1312"))

}
