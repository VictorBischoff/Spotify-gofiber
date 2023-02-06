package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/victorbischoff/GOFIBER-TMPL/database"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/models"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Fiber-template",
	})
}

func GetSongs(c *fiber.Ctx) error {
	db := database.DBConn
	var songs []models.Song
	db.Find(&songs)
	return c.JSON(songs)
}

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

func AddSong(c *fiber.Ctx) error {
	db := database.DBConn
	var song models.Song

	song.Title = "CREAM"
	song.Rating = 5
	song.Artist = "Wu-tang-clan"
	db.Create(&song)
	return c.JSON(song)
}

func UpdateSong(c *fiber.Ctx) error {
	return c.SendString("UPDATE SONGS")
}

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
