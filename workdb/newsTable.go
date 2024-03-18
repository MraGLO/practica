package workdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5"
)

func InsertNews(news *model.News, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "INSERT INTO news(shortname, body, author, published, body_full) VALUES($1, $2, $3, $4, $5)",
		news.Shortname, news.Body, news.Author, false, news.BodyFull); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeNews(id int, news *model.News, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "UPDATE category SET shortname = $1, body = $2, author = $3, changed_time = $4 published = $5, body_full = $6 WHERE id = $7",
		news.Shortname, news.Body, news.Author, time.Now(), news.Published, news.BodyFull, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNews(id int, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "DELETE FROM news_category WHERE news_id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	if _, err := conn.Exec(context.Background(), "DELETE FROM news_tag WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	if _, err := conn.Exec(context.Background(), "DELETE FROM news WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNews(conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news"); err != nil {
		return "news", err
	} else {
		defer rows.Close()
		var tmp model.News
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.Shortname, &tmp.Body, &tmp.Author, &tmp.PublishedTime, &tmp.ChangedTime, &tmp.Published, &tmp.BodyFull)
			message = message + fmt.Sprintf("ID: %d   shortname: %s   body: %s   author: %s   published_time: %s   changed_time: %s   published: %t   body_full: %s\n",
				tmp.ID, tmp.Shortname, tmp.Body, tmp.Author, tmp.PublishedTime.GoString(), tmp.ChangedTime.GoString(), tmp.Published, tmp.BodyFull)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return "message", err

		}
	}
	return message, nil
}

func SelectAnNews(id int, conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, shortname, body, author, published_time, changed_time, published, body_full  FROM news WHERE id = $1", id); err != nil {
		return "news", err
	} else {
		defer rows.Close()
		var tmp model.News
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.Shortname, &tmp.Body, &tmp.Author, &tmp.PublishedTime, &tmp.ChangedTime, &tmp.Published, &tmp.BodyFull)
			message = message + fmt.Sprintf("ID: %d   shortname: %s   body: %s   author: %s   published_time: %s   changed_time: %s   published: %t   body_full: %s\n",
				tmp.ID, tmp.Shortname, tmp.Body, tmp.Author, tmp.PublishedTime.GoString(), tmp.ChangedTime.GoString(), tmp.Published, tmp.BodyFull)
		}
		if rows.Err() != nil {
			log.Println(err)
			return "message", err

		}

	}
	return message, nil
}
