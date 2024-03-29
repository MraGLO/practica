package model

type NewNews struct {
	Shortname     string `json:"shortname"`
	Body          string `json:"body"`
	Author        string `json:"author"`
	PublishedTime string `json:"publishedTime"`
	BodyFull      string `json:"bodyFull"`
	Categories    []int  `json:"category"`
	Tags          []int  `json:"tag"`
}
