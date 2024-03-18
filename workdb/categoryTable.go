package workdb

import (
	"context"
	"fmt"
	"log"

	"github.com/MraGLO/practica/model"
	"github.com/jackc/pgx/v5"
)

func InsertCategory(category *model.Category, conn *pgx.Conn) error {

	if _, err := conn.Exec(context.Background(), "INSERT INTO category(name, name_en) VALUES($1, $2)", category.CategoryName, category.CategoryNameEN); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeCategory(id int, category *model.Category, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "UPDATE category SET name = $1, name_en = $2 WHERE id = $3", category.CategoryName, category.CategoryNameEN, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCategory(id int, conn *pgx.Conn) error {
	if _, err := conn.Exec(context.Background(), "DELETE FROM category WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllCategory(conn *pgx.Conn) (string, error) {
	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, name, name_en  FROM category"); err != nil {
		return "categories", err
	} else {
		defer rows.Close()
		var tmp model.Category
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.CategoryName, &tmp.CategoryNameEN)
			message = message + fmt.Sprintf("ID: %d   CategoryName: %s   CategoryNameEN: %s \n", tmp.ID, tmp.CategoryName, tmp.CategoryNameEN)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return "message", err

		}
	}
	return message, nil
}

func SelectAnCategory(id int, conn *pgx.Conn) (string, error) {

	message := ""
	if rows, err := conn.Query(context.Background(), "SELECT id, name, name_en FROM category WHERE id = $1", id); err != nil {
		return "category", err
	} else {
		defer rows.Close()
		var tmp model.Category
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.CategoryName, &tmp.CategoryNameEN)
			message = fmt.Sprintf("ID: %d   CategoryName: %s   CategoryNameEN: %s", tmp.ID, tmp.CategoryName, tmp.CategoryNameEN)
		}
		if rows.Err() != nil {
			log.Println(err)
			return "message", err

		}

	}

	return message, nil

}
