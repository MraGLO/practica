package workdb

import (
	"context"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertTags(tags *model.Tag, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "INSERT INTO tags(name, name_en) VALUES($1, $2)", tags.TagName, tags.TagNameEN); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeTags(id int, tags *model.Tag, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "UPDATE tags SET name = $1, name_en = $2 WHERE id = $3", tags.TagName, tags.TagNameEN, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteTags(id int, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllTags(dbpool *pgxpool.Pool) ([]model.Tag, error) {
	var tags []model.Tag
	if rows, err := dbpool.Query(context.Background(), "SELECT id, name, name_en  FROM tags"); err != nil {
		log.Println(err)
		return tags, err
	} else {
		defer rows.Close()
		var tmp model.Tag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.TagName, &tmp.TagNameEN)
			tags = append(tags, tmp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return tags, err

		}
	}
	return tags, nil
}

func SelectAnTags(id int, dbpool *pgxpool.Pool) (model.Tag, error) {
	var tmp model.Tag
	if rows, err := dbpool.Query(context.Background(), "SELECT id, name, name_en FROM tags WHERE id = $1", id); err != nil {
		log.Println(err)
		return tmp, err
	} else {
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.TagName, &tmp.TagNameEN)
		}
		if rows.Err() != nil {
			log.Println(err)
			return tmp, err

		}

	}
	return tmp, nil
}
