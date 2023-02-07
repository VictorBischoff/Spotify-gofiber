package models

import (
	"gorm.io/gorm"
)

// Song contians the columns that the song table will have in postgres
// the gorm.Model struct inside of the Song struct, contains useful columns for a table.
type Song struct {
	gorm.Model
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Length    float32 `json:"length"`
	Available bool    `json:"available"`
	Rating    int     `json:"rating"`
}
