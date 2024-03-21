package model

import (
	"time"
)

type News struct {
	ID            int       `json:"id"`
	Shortname     string    `json:"shortname"`
	Body          string    `json:"body"`
	Author        string    `json:"author"`
	PublishedTime time.Time `json:"publishedTime"`
	ChangedTime   time.Time `json:"changedTime"`
	Published     bool      `json:"published"`
	BodyFull      string    `json:"bodyFull"`
}
