package repository

import (
	"context"

	"github.com/MraGLO/practica/pkg/model"
)

func (d *DatabaseRepo) SelectAllCategories() (categories []model.Category, err error) {
	var tmp model.Category
	rows, err := d.db.Query(context.Background(), "SELECT id, name, name_en  FROM category")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.CategoryName, &tmp.CategoryNameEN)
		if err != nil {
			return
		}
		categories = append(categories, tmp)

	}

	return
}

func (d *DatabaseRepo) SelectCategoryByID(id int) (category model.Category, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, name, name_en FROM category WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&category.ID, &category.CategoryName, &category.CategoryNameEN)
		if err != nil {
			return
		}
	}
	return

}

func (d *DatabaseRepo) SelectAllTags() (tags []model.Tag, err error) {
	var tmp model.Tag
	rows, err := d.db.Query(context.Background(), "SELECT id, name, name_en  FROM tags")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.TagName, &tmp.TagNameEN)
		if err != nil {
			return
		}
		tags = append(tags, tmp)

	}
	return
}

func (d *DatabaseRepo) SelectTagByID(id int) (tag model.Tag, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, name, name_en FROM tags WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tag.ID, &tag.TagName, &tag.TagNameEN)
		if err != nil {
			return
		}
	}
	return

}

func (d *DatabaseRepo) SelectAllNewsCategories() (newsCategories []model.NewsCategory, err error) {
	var tmp model.NewsCategory
	rows, err := d.db.Query(context.Background(), "SELECT id, news_id, category_id  FROM news_category")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.CategoryID)
		if err != nil {
			return
		}
		newsCategories = append(newsCategories, tmp)

	}
	return
}

func (d *DatabaseRepo) SelectNewsCategoryByID(id int) (newsCategory model.NewsCategory, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, news_id, category_id  FROM news_category WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&newsCategory.ID, &newsCategory.NewsID, &newsCategory.CategoryID)
		if err != nil {
			return
		}

	}
	return
}

func (d *DatabaseRepo) SelectAllNewsTags() (newsTags []model.NewsTag, err error) {
	var tmp model.NewsTag
	rows, err := d.db.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_tag")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.TagID)
		if err != nil {
			return
		}
		newsTags = append(newsTags, tmp)

	}
	return
}

func (d *DatabaseRepo) SelectNewsTagByID(id int) (newsTag model.NewsTag, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_tag WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&newsTag.ID, &newsTag.NewsID, &newsTag.TagID)
		if err != nil {
			return
		}
	}
	return
}

func (d *DatabaseRepo) SelectAllNews() (news []model.News, err error) {
	var tmp model.News
	rows, err := d.db.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&tmp.ID, &tmp.Shortname, &tmp.Body, &tmp.Author, &tmp.PublishedTime, &tmp.ChangedTime, &tmp.Published, &tmp.BodyFull)
		if err != nil {
			return
		}
		news = append(news, tmp)
	}
	return
}

func (d *DatabaseRepo) SelectNewsByID(id int) (news model.News, err error) {
	rows, err := d.db.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news WHERE id = $1", id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&news.ID, &news.Shortname, &news.Body, &news.Author, &news.PublishedTime, &news.ChangedTime, &news.Published, &news.BodyFull)
		if err != nil {
			return
		}
	}
	return
}
