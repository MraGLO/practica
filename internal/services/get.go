package services

import "github.com/MraGLO/practica/pkg/model"

func (s *Services) GetAllCategories() ([]model.Category, error) {
	return s.Database.db.SelectAllCategories()
}

func (s *Services) GetAllTags() ([]model.Tag, error) {
	return s.Database.db.SelectAllTags()
}

func (s *Services) GetAllNewsCategories() ([]model.NewsCategory, error) {
	return s.Database.db.SelectAllNewsCategories()
}

func (s *Services) GetAllNewsTags() ([]model.NewsTag, error) {
	return s.Database.db.SelectAllNewsTags()
}

func (s *Services) GetAllNews() ([]model.News, error) {
	return s.Database.db.SelectAllNews()
}

func (s *Services) GetCategoryByID(id int) (model.Category, error) {
	return s.Database.db.SelectCategoryByID(id)
}

func (s *Services) GetTagByID(id int) (model.Tag, error) {
	return s.Database.db.SelectTagByID(id)
}

func (s *Services) GetNewsCategoryByID(id int) (model.NewsCategory, error) {
	return s.Database.db.SelectNewsCategoryByID(id)
}

func (s *Services) GetNewsTagByID(id int) (model.NewsTag, error) {
	return s.Database.db.SelectNewsTagByID(id)
}

func (s *Services) GetNewsByID(id int) (model.News, error) {
	return s.Database.db.SelectNewsByID(id)
}
