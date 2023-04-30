package models

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Year        int       `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	MPPARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
