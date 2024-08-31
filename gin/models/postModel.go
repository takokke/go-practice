package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ImageURL     string
	CreatureName string
	Caption      string
	Address      string
	Latitude     float64
	Longitude    float64
	UserID       uint
	User         User
}
