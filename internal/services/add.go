package services

import "github.com/MraGLO/practica/pkg/model"

func (s *Services) AddCategory(category model.Category) error {
	return s.Database.db.InsertCategory(&category)
}

func (s *Services) AddTags(tags model.Tag) error {
	return s.Database.db.InsertTags(&tags)
}

func (s *Services) AddNewsCategory(newsCategory model.NewsCategory) error {
	return s.Database.db.InsertNewsCategory(&newsCategory)
}

func (s *Services) AddNewsTag(newsTag model.NewsTag) error {
	return s.Database.db.InsertNewsTag(&newsTag)
}

func (s *Services) AddNews(news model.NewNews, isLenCategories bool, isLenTags bool) error {
	return s.Database.db.InsertNews(&news, isLenCategories, isLenTags)
}
