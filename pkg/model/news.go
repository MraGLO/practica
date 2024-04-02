package model

type News struct {
	ID               int         `json:"id"`
	Shortname        string      `json:"shortname"`
	Body             string      `json:"body"`
	Author           string      `json:"author"`
	PublishedTime    interface{} `json:"publishedTime"`
	ChangedTime      interface{} `json:"changedTime"`
	Published        bool        `json:"published"`
	BodyFull         interface{} `json:"bodyFull"`
	TopicImageSrc    interface{} `json:"topicImageSrc"`
	TopicImageSrcSet interface{} `json:"topicImageSrcSet"`
}
