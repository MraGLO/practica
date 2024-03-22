package repository

import (
	"context"
	"time"

	"github.com/MraGLO/practica/pkg/model"
)

func (d *DatabaseRepo) UpdateCategory(id int, category *model.Category) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE category SET name = $1, name_en = $2 WHERE id = $3", category.CategoryName, category.CategoryNameEN, id)
	return
}

func (d *DatabaseRepo) UpdateTag(id int, tags *model.Tag) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE tags SET name = $1, name_en = $2 WHERE id = $3", tags.TagName, tags.TagNameEN, id)
	return
}

func (d *DatabaseRepo) UpdateNewsCategory(id int, newsCategory *model.NewsCategory) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE news_category SET news_id = $1, category_id = $2 WHERE id = $3", newsCategory.NewsID, newsCategory.CategoryID, id)
	return
}

func (d *DatabaseRepo) UpdateNewsTag(id int, newsTag *model.NewsTag) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE news_tag SET news_id = $1, tag_id = $2 WHERE id = $3", newsTag.NewsID, newsTag.TagID, id)
	return
}

func (d *DatabaseRepo) UpdateNews(id int, news *model.News) (err error) {
	_, err = d.db.Exec(context.Background(), "UPDATE category SET shortname = $1, body = $2, author = $3, changed_time = $4 published = $5, body_full = $6 WHERE id = $7",
		news.Shortname, news.Body, news.Author, time.Now(), news.Published, news.BodyFull, id)
	return
}
