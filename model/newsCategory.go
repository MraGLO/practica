package model

type NewsCategory struct {
	ID         int `json: "id"`
	NewsID     int `json: "newsID"`
	CategoryID int `json: "categoryID"`
}
