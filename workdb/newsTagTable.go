package workdb

import (
	"context"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewsTag(newsTag *model.NewsTag, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "INSERT INTO news_tag(news_id, tag_id) VALUES($1, $2)", newsTag.NewsID, newsTag.TagID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeNewsTag(id int, newsTag *model.NewsTag, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "UPDATE news_tag SET news_id = $1, tag_id = $2 WHERE id = $3", newsTag.NewsID, newsTag.TagID, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNewsTag(id int, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "DELETE FROM  news_tag WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNewsTag(dbpool *pgxpool.Pool) ([]model.NewsTag, error) {
	var newsTags []model.NewsTag
	if rows, err := dbpool.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_tag"); err != nil {
		return newsTags, err
	} else {
		defer rows.Close()
		var tmp model.NewsTag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.TagID)
			newsTags = append(newsTags, tmp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return newsTags, err

		}
	}
	return newsTags, nil
}

func SelectAnNewsTag(id int, dbpool *pgxpool.Pool) (model.NewsTag, error) {
	var tmp model.NewsTag
	if rows, err := dbpool.Query(context.Background(), "SSELECT id, news_id, tag_id  FROM news_tag WHERE id = $1", id); err != nil {
		return tmp, err
	} else {
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.TagID)
		}
		if rows.Err() != nil {
			log.Println(err)
			return tmp, err

		}

	}
	return tmp, nil
}
