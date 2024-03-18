package workdb

import (
	"context"
	"fmt"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5"
)

func InsertTags(tags *model.Tag, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "INSERT INTO tags(name, name_en) VALUES($1, $2)", tags.TagName, tags.TagNameEN); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeTags(id int, tags *model.Tag, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "UPDATE tags SET name = $1, name_en = $2 WHERE id = $3", tags.TagName, tags.TagNameEN, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteTags(id int, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllTags(conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, name, name_en  FROM tags"); err != nil {
		log.Println(err)
		return "categories", err
	} else {
		defer rows.Close()
		var tmp model.Tag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.TagName, &tmp.TagNameEN)
			message = message + fmt.Sprintf("ID: %d   CategoryName: %s   CategoryNameEN: %s \n", tmp.ID, tmp.TagName, tmp.TagNameEN)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return "message", err

		}
	}
	return message, nil
}

func SelectAnTags(id int, conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, name, name_en FROM tags WHERE id = $1", id); err != nil {
		log.Println(err)
		return "category", err
	} else {
		defer rows.Close()
		var tmp model.Tag
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.TagName, &tmp.TagNameEN)
			message = fmt.Sprintf("ID: %d   CategoryName: %s   CategoryNameEN: %s", tmp.ID, tmp.TagName, tmp.TagNameEN)
		}
		if rows.Err() != nil {
			log.Println(err)
			return "message", err

		}

	}
	return message, nil
}
