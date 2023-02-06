package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Title     string  `json:"title"`
	Artist    string  `json:"artist"`
	Length    float32 `json:"length"`
	Available bool    `json:"available"`
	Rating    int     `json:"rating"`
}
