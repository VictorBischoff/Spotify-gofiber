package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"

	"github.com/victorbischoff/GOFIBER-TMPL/database"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/config"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/routers"
)

// Server creates Fiber instance and initializes the database connection.
// Server takes a string as a parameter, which will be the port that fiber will listen on.
func Server(port string) {
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	

	routers.AllRouters(app)

	log.Fatal(app.Listen(port))

}
