package model

type Tag struct {
	ID        int    `json: "id"`
	TagName   string `json: "tagName"`
	TagNameEN string `json: "tagNameEN"`
}
