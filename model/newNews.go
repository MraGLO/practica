package model

type NewNews struct {
	Shortname  string `json:"shortname"`
	Body       string `json:"body"`
	Author     string `json:"author"`
	BodyFull   string `json:"bodyFull"`
	Categories []int  `json:"categories"`
	Tags       []int  `json:"tags"`
}
