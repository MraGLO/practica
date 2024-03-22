package workdb

import (
	"context"
	"log"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InsertCategory(category *model.Category, dbpool *pgxpool.Pool) error {

	if _, err := dbpool.Exec(context.Background(), "INSERT INTO category(name, name_en) VALUES($1, $2)", category.CategoryName, category.CategoryNameEN); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ChangeCategory(id int, category *model.Category, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "UPDATE category SET name = $1, name_en = $2 WHERE id = $3", category.CategoryName, category.CategoryNameEN, id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteCategory(id int, dbpool *pgxpool.Pool) error {
	if _, err := dbpool.Exec(context.Background(), "DELETE FROM category WHERE id = $1", id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SelectAllCategory(dbpool *pgxpool.Pool) ([]model.Category, error) {
	var categories []model.Category
	if rows, err := dbpool.Query(context.Background(), "SELECT id, name, name_en  FROM category"); err != nil {
		return categories, err
	} else {
		defer rows.Close()
		var tmp model.Category
		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.CategoryName, &tmp.CategoryNameEN)
			categories = append(categories, tmp)
		}
		if rows.Err() != nil {
			log.Println(rows.Err())
			return categories, err

		}
	}
	return categories, nil
}

func SelectAnCategory(id int, dbpool *pgxpool.Pool) (model.Category, error) {
	var tmp model.Category
	if rows, err := dbpool.Query(context.Background(), "SELECT id, name, name_en FROM category WHERE id = $1", id); err != nil {
		return tmp, err
	} else {
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&tmp.ID, &tmp.CategoryName, &tmp.CategoryNameEN)
		}
		if rows.Err() != nil {
			log.Println(err)
			return tmp, err

		}

	}

	return tmp, nil

}
