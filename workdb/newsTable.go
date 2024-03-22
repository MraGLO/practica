package workdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNews(news *model.NewNews, dbpool *pgxpool.Pool) error {
	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
			log.Println(err)
			fmt.Println(" NOT OK")
		} else {
			tx.Commit(context.Background())
			fmt.Println("ok")
		}
	}()

	var id int

	if row, err := tx.Query(context.Background(), "INSERT INTO news(shortname, body, author, published, body_full) VALUES($1, $2, $3, $4, $5) RETURNING id",
		news.Shortname, news.Body, news.Author, true, news.BodyFull); err != nil {
		log.Println(err)
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
			log.Println(err)
			return err
		}
	}

	for _, tagID := range news.Tags {
		if _, err := tx.Exec(context.Background(), "INSERT INTO news_tag(news_id, tag_id) VALUES($1, $2)",
			id, tagID); err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func ChangeNews(id int, news *model.News, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "UPDATE category SET shortname = $1, body = $2, author = $3, changed_time = $4 published = $5, body_full = $6 WHERE id = $7",
		news.Shortname, news.Body, news.Author, time.Now(), news.Published, news.BodyFull, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNews(id int, dbpool *pgxpool.Pool) error {
	tx, err := dbpool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(context.Background())
			log.Println(err)
			fmt.Println(" NOT OK")
		} else {
			tx.Commit(context.Background())
			fmt.Println("ok")
		}
	}()

	if _, err := tx.Exec(context.Background(), "DELETE FROM news_category WHERE news_id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	if _, err := tx.Exec(context.Background(), "DELETE FROM news_tag WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	if _, err := tx.Exec(context.Background(), "DELETE FROM news WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNews(dbpool *pgxpool.Pool) ([]model.News, error) {
	var news []model.News
	if rows, err := dbpool.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news"); err != nil {
		return news, err
	} else {
		defer rows.Close()
		var tmp model.News
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.Shortname, &tmp.Body, &tmp.Author, &tmp.PublishedTime, &tmp.ChangedTime, &tmp.Published, &tmp.BodyFull)
			news = append(news, tmp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return news, err

		}
	}
	return news, nil
}

func SelectAnNews(id int, dbpool *pgxpool.Pool) (model.News, error) {
	var tmp model.News
	if rows, err := dbpool.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news WHERE id = $1", id); err != nil {
		return tmp, err
	} else {
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.Shortname, &tmp.Body, &tmp.Author, &tmp.PublishedTime, &tmp.ChangedTime, &tmp.Published, &tmp.BodyFull)
		}
		if rows.Err() != nil {
			log.Println(err)
			return tmp, err

		}

	}
	return tmp, nil
}
