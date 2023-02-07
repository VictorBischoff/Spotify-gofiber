package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/victorbischoff/GOFIBER-TMPL/database"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/config"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/routers"
)

func main() {
	config, err := config.LoadConfig("./app.env")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	database.InitDatabase(&config)

	engine := html.New("./public", ".gohtml")
	engine.Reload(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public", fiber.Static{
		Compress: true,
	})

	

	routers.Routers(app)

	log.Fatal(app.Listen(":1312"))

}
