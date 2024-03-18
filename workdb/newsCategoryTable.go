package workdb

import (
	"context"
	"fmt"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5"
)

func InsertNewsCategory(newsCategory *model.NewsCategory, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "INSERT INTO news_category(news_id, tag_id) VALUES($1, $2)", newsCategory.NewsID, newsCategory.CategoryID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeNewsCategory(id int, newsCategory *model.NewsCategory, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "UPDATE news_category SET news_id = $1, tag_id = $2 WHERE id = $3", newsCategory.NewsID, newsCategory.CategoryID, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNewsCategory(id int, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "DELETE FROM  news_category WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNewsCategory(conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_category"); err != nil {
		return "NewsCategories", err
	} else {
		defer rows.Close()
		var tmp model.NewsCategory
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.CategoryID)
			message = message + fmt.Sprintf("ID: %d   NewsID: %d   CategoryID: %d \n", tmp.ID, tmp.NewsID, tmp.CategoryID)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return "message", err

		}
	}
	return message, nil
}

func SelectAnNewsCategory(id int, conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_category WHERE id = $1", id); err != nil {
		return "NewsCategories", err
	} else {
		defer rows.Close()
		var tmp model.NewsCategory
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.CategoryID)
			message = message + fmt.Sprintf("ID: %d   NewsID: %d   CategoryID: %d \n", tmp.ID, tmp.NewsID, tmp.CategoryID)
		}
		if rows.Err() != nil {
			log.Println(err)
			return "message", err

		}

	}
	return message, nil
}
