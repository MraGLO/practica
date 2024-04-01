package repository

import (
	"github.com/MraGLO/practica/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database interface {
	SelectAllCategories() (categories []model.Category, err error)
	SelectCategoryByID(id int) (category model.Category, err error)
	SelectAllTags() (tags []model.Tag, err error)
	SelectTagByID(id int) (tag model.Tag, err error)
	SelectAllNewsCategories() (newsCategories []model.NewsCategory, err error)
	SelectNewsCategoryByID(id int) (newsCategory model.NewsCategory, err error)
	SelectAllNewsTags() (newsTags []model.NewsTag, err error)
	SelectNewsTagByID(id int) (newsTag model.NewsTag, err error)
	SelectAllNews() (news []model.News, err error)
	SelectNewsByID(id int) (news model.News, err error)

	InsertCategory(category *model.Category) (err error)
	InsertTags(tags *model.Tag) (err error)
	InsertNewsCategory(newsCategory *model.NewsCategory) (err error)
	InsertNewsTag(newsTag *model.NewsTag) (err error)
	InsertNews(news *model.NewNews, isLenCategories bool, isLenTags bool) (err error)

	UpdateCategory(id int, category *model.Category) (err error)
	UpdateTag(id int, tags *model.Tag) (err error)
	UpdateNewsCategory(id int, newsCategory *model.NewsCategory) (err error)
	UpdateNewsTag(id int, newsTag *model.NewsTag) (err error)
	UpdateNews(id int, news *model.News) (err error)

	DeleteCategory(id int) (count int, err error)
	DeleteTag(id int) (count int, err error)
	DeleteNewsCategory(id int) (count int, err error)
	DeleteNewsTag(id int) (count int, err error)
	DeleteNews(id int) (count int, err error)
}

type Repository struct {
	Database Database
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{Database: newDatabaseRepo(db)}
}
