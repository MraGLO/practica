package model

type NewsTag struct {
	ID     int `json: "id"`
	NewsID int `json: "newsID"`
	TagID  int `json: "tagID"`
}
