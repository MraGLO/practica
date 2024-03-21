package repository

import (
	"context"
)

func (d *DatabaseRepo) DeleteCategory(id int) (err error) {
	_, err = d.db.Exec(context.Background(), "DELETE FROM category WHERE id = $1", id)
	return
}

func (d *DatabaseRepo) DeleteTags(id int) (err error) {
	_, err = d.db.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", id)
	return
}

func (d *DatabaseRepo) DeleteNewsCategory(id int) (err error) {
	_, err = d.db.Exec(context.Background(), "DELETE FROM  news_category WHERE id = $1", id)
	return
}

func (d *DatabaseRepo) DeleteNewsTag(id int) (err error) {
	_, err = d.db.Exec(context.Background(), "DELETE FROM  news_tag WHERE id = $1", id)
	return
}

func (d *DatabaseRepo) DeleteNews(id int) (err error) {
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

	if _, err = tx.Exec(context.Background(), "DELETE FROM news_category WHERE news_id = $1", id); err != nil {
		return
	}
	if _, err = tx.Exec(context.Background(), "DELETE FROM news_tag WHERE id = $1", id); err != nil {
		return
	}
	if _, err = tx.Exec(context.Background(), "DELETE FROM news WHERE id = $1", id); err != nil {
		return
	}
	return
}
