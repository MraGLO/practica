package workdb

import (
	"context"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertNewsCategory(newsCategory *model.NewsCategory, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "INSERT INTO news_category(news_id, tag_id) VALUES($1, $2)", newsCategory.NewsID, newsCategory.CategoryID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeNewsCategory(id int, newsCategory *model.NewsCategory, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "UPDATE news_category SET news_id = $1, tag_id = $2 WHERE id = $3", newsCategory.NewsID, newsCategory.CategoryID, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteNewsCategory(id int, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "DELETE FROM  news_category WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllNewsCategory(dbpool *pgxpool.Pool) ([]model.NewsCategory, error) {
	var newCategories []model.NewsCategory
	if rows, err := dbpool.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_category"); err != nil {
		return newCategories, err
	} else {
		defer rows.Close()
		var tmp model.NewsCategory
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.CategoryID)
			newCategories = append(newCategories, tmp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return newCategories, err

		}
	}
	return newCategories, nil
}

func SelectAnNewsCategory(id int, dbpool *pgxpool.Pool) (model.NewsCategory, error) {
	var tmp model.NewsCategory
	if rows, err := dbpool.Query(context.Background(), "SELECT id, news_id, tag_id  FROM news_category WHERE id = $1", id); err != nil {
		return tmp, err
	} else {
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.NewsID, &tmp.CategoryID)
		}
		if rows.Err() != nil {
			log.Println(err)
			return tmp, err

		}

	}
	return tmp, nil
}
