package services

import "github.com/MraGLO/practica/pkg/model"

func (s *Services) UpdateCategory(id int, category model.Category) error {
	return s.Database.db.UpdateCategory(id, &category)
}

func (s *Services) UpdateTags(id int, tags model.Tag) error {
	return s.Database.db.UpdateTag(id, &tags)
}

func (s *Services) UpdateNewsCategory(id int, newsCategory model.NewsCategory) error {
	return s.Database.db.UpdateNewsCategory(id, &newsCategory)
}

func (s *Services) UpdateNewsTag(id int, newsTag model.NewsTag) error {
	return s.Database.db.UpdateNewsTag(id, &newsTag)
}

func (s *Services) UpdateNews(id int, news model.News) error {
	return s.Database.db.UpdateNews(id, &news)
}
