package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/victorbischoff/GOFIBER-TMPL/database"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/models"
)

// Indexhandler is for healthchecking the api.
func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Fiber-template",
	})
}

// GetSongs returns all the songs that are stored in the database in json format.
func GetSongs(c *fiber.Ctx) error {
	db := database.DBConn
	var songs []models.Song
	db.Find(&songs)
	return c.JSON(songs)
}

// GetSong returns one song using the id that is stored in the database table in json format.
func GetSong(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	db := database.DBConn
	var song models.Song
	db.Find(&song, id)
	return c.JSON(song)
}

// AddSong allows the client to add a new song to the database, and return the result in json format.
func AddSong(c *fiber.Ctx) error {
	db := database.DBConn
	var song *models.Song = &models.Song{}

	if err := c.BodyParser(song); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}

	db.Create(&song)
	return c.JSON(song)
}

//TODO
func UpdateSong(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var song models.Song

	db.First(&song, id)

	if song.Title == "" {
		c.Status(500).SendString("no song with the id")
		return db.Error
	}

	db.Update("wu", &song)
	return c.JSON("song updated")
}

// DeleteSong will delete a song by using the id 
func DeleteSong(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var song models.Song

	db.First(&song, id)
	if song.Title == "" {
		c.Status(500).SendString("no song with the id")
		return db.Error
	}

	db.Delete(&song)
	return c.JSON("Song deleted")

}
