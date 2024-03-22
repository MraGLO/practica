package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

func (d *DatabaseRepo) DeleteCategory(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM category WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteTag(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteNewsCategory(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM  news_category WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteNewsTag(id int) (count int, err error) {
	com, err := d.db.Exec(context.Background(), "DELETE FROM  news_tag WHERE id = $1", id)
	if err != nil {
		return
	}

	count = int(com.RowsAffected())
	return
}

func (d *DatabaseRepo) DeleteNews(id int) (count int, err error) {
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
	var com pgconn.CommandTag
	if com, err = tx.Exec(context.Background(), "DELETE FROM news_category WHERE news_id = $1", id); err != nil {
		if err != nil {
			return
		}

		count = int(com.RowsAffected())

	}
	if _, err = tx.Exec(context.Background(), "DELETE FROM news_tag WHERE id = $1", id); err != nil {
		return
	}
	if _, err = tx.Exec(context.Background(), "DELETE FROM news WHERE id = $1", id); err != nil {
		return
	}
	return
}
