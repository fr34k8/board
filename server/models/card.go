package models

import (
	"time"
)

type Card struct {
	Id uint `json:"id"`
	Project uint `json:"project"`
	Number uint `json:"number"`
	Title string `json:"title"`
	Description string `json:"title"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
