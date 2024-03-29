package repository

import (
	"context"

	"github.com/MraGLO/practica/pkg/model"
)

func (d *DatabaseRepo) InsertCategory(category *model.Category) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO category(name, name_en) VALUES($1, $2)", category.CategoryName, category.CategoryNameEN)
	return
}

func (d *DatabaseRepo) InsertTags(tags *model.Tag) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO tags(name, name_en) VALUES($1, $2)", tags.TagName, tags.TagNameEN)
	return
}

func (d *DatabaseRepo) InsertNewsCategory(newsCategory *model.NewsCategory) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO news_category(news_id, category_id) VALUES($1, $2)", newsCategory.NewsID, newsCategory.CategoryID)
	return
}

func (d *DatabaseRepo) InsertNewsTag(newsTag *model.NewsTag) (err error) {
	_, err = d.db.Exec(context.Background(), "INSERT INTO news_tag(news_id, tag_id) VALUES($1, $2)", newsTag.NewsID, newsTag.TagID)
	return
}

func (d *DatabaseRepo) InsertNews(news *model.NewNews) (err error) {
	tx, err := d.db.Begin(context.Background())
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
			return
		} else {
			tx.Commit(context.Background())
			return
		}
	}()

	var id int
	if row, err := tx.Query(context.Background(), "INSERT INTO news(shortname, body, author, published, body_full) VALUES($1, $2, $3, $4, $5) RETURNING id",
		news.Shortname, news.Body, news.Author, true, news.BodyFull); err != nil {
		return err
	} else {
		defer row.Close()
		for row.Next() {
			row.Scan(&id)
		}
	}
	for _, categoryID := range news.Categories {
		if _, err := tx.Exec(context.Background(), "INSERT INTO news_category(news_id, category_id) VALUES($1, $2)",
			id, categoryID); err != nil {
			return err
		}
	}

	for _, tagID := range news.Tags {
		if _, err := tx.Exec(context.Background(), "INSERT INTO news_tag(news_id, tag_id) VALUES($1, $2)",
			id, tagID); err != nil {
			return err
		}
	}
	return
}
