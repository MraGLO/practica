package model

type Category struct {
	ID             int    `json: "id"`
	CategoryName   string `json: "categoryName"`
	CategoryNameEN string `json: "categoryNameEN"`
}
