package workdb

import (
	"context"
	"fmt"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5"
)

func InsertNewsTag(newsTag *model.NewsTag, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "INSERT INTO news_tag(news_id, tag_id) VALUES($1, $2)", newsTag.NewsID, newsTag.TagID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeNewsTag(id int, newsTag *model.NewsTag, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "UPDATE news_tag SET news_id = $1, tag_id = $2 WHERE id = $3", newsTag.NewsID, newsTag.TagID, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNewsTag(id int, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "DELETE FROM  news_tag WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNewsTag(conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_tag"); err != nil {
		return "NewsTag", err
	} else {
		defer rows.Close()
		var tmp model.NewsTag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.TagID)
			message = message + fmt.Sprintf("ID: %d   NewsID: %d   TagID: %d \n", tmp.ID, tmp.NewsID, tmp.TagID)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return "message", err

		}
	}
	return message, nil
}

func SelectAnNewsTag(id int, conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SSELECT id, news_id, tag_id  FROM news_tag WHERE id = $1", id); err != nil {
		return "NewsCategories", err
	} else {
		defer rows.Close()
		var tmp model.NewsTag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.TagID)
			message = message + fmt.Sprintf("ID: %d   NewsID: %d   TagID: %d \n", tmp.ID, tmp.NewsID, tmp.TagID)
		}
		if rows.Err() != nil {
			log.Println(err)
			return "message", err

		}

	}
	return message, nil
}
