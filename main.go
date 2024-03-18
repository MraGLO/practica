package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/MraGLO/practica/model"
	"github.com/MraGLO/practica/workdb"
)

func main() {

	conn := workdb.ConnectDB()

	defer conn.Close(context.Background())

	app := fiber.New()
	app.Static("/", "./PublicationPage")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("main")
	})

	// ------- /category --------------------------------------------

	app.Get("/category", func(c *fiber.Ctx) error {
		message, err := workdb.SelectAllCategory(conn)

		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(message)

	})
	app.Post("/category", func(c *fiber.Ctx) error {
		newCategory := new(model.Category)
		if err := c.BodyParser(newCategory); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertCategory(newCategory, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString("Новая категория была добавлена")
	})

	app.Get("/category/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		category, err := workdb.SelectAnCategory(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(category)
	})

	app.Put("/category/:id", func(c *fiber.Ctx) error {
		newCategory := new(model.Category)
		if err := c.BodyParser(newCategory); err != nil {
			log.Println(err)
			return err
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.ChangeCategory(id, newCategory, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Категория с id = %d была добавлена", id)
		return c.SendString(message)

	})
	app.Delete("/category/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.DeleteCategory(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Категория с id = %d была удалена", id)
		return c.SendString(message)
	})

	// ------- /category --------------------------------------------

	// ------- /tag --------------------------------------------

	app.Get("/tag", func(c *fiber.Ctx) error {
		message, err := workdb.SelectAllTags(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(message)

	})

	app.Post("/tag", func(c *fiber.Ctx) error {
		newTag := new(model.Tag)
		if err := c.BodyParser(newTag); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertTags(newTag, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString("Новый тег был добавлен")
	})

	app.Get("/tag/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		tag, err := workdb.SelectAnTags(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(tag)
	})

	app.Put("/tag/:id", func(c *fiber.Ctx) error {
		newTag := new(model.Tag)
		if err := c.BodyParser(newTag); err != nil {
			log.Println(err)
			return err
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.ChangeTags(id, newTag, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Тег с id = %d был изменен", id)
		return c.SendString(message)

	})

	app.Delete("/tag/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.DeleteTags(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Тег с id = %d был удален", id)
		return c.SendString(message)
	})

	// ------- /tags --------------------------------------------
	//
	//
	//
	// ------- /newsCategory --------------------------------------------
	app.Get("/newsCategory", func(c *fiber.Ctx) error {
		message, err := workdb.SelectAllNewsCategory(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(message)

	})

	app.Post("/newsCategory", func(c *fiber.Ctx) error {
		newNewsCategory := new(model.NewsCategory)
		if err := c.BodyParser(newNewsCategory); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNewsCategory(newNewsCategory, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString("Новая запись была добавлена")
	})

	app.Get("/newsCategory/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		tag, err := workdb.SelectAnNewsCategory(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(tag)
	})

	app.Put("/newsCategory/:id", func(c *fiber.Ctx) error {
		newNewsCategory := new(model.NewsCategory)
		if err := c.BodyParser(newNewsCategory); err != nil {
			log.Println(err)
			return err
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.ChangeNewsCategory(id, newNewsCategory, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была изменена", id)
		return c.SendString(message)

	})

	app.Delete("/newsCategory/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.DeleteNewsCategory(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была удалена", id)
		return c.SendString(message)
	})
	// ------- /newsCategory --------------------------------------------
	//
	//
	//
	// ------- /newsTag --------------------------------------------
	app.Get("/newsTag", func(c *fiber.Ctx) error {
		message, err := workdb.SelectAllNewsTag(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(message)

	})

	app.Post("/newsTag", func(c *fiber.Ctx) error {
		newNewsTag := new(model.NewsTag)
		if err := c.BodyParser(newNewsTag); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNewsTag(newNewsTag, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString("Новая запись была добавлена")
	})

	app.Get("/newsTag/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		tag, err := workdb.SelectAnNewsTag(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(tag)
	})

	app.Put("/newsTag/:id", func(c *fiber.Ctx) error {
		newNewsTag := new(model.NewsTag)
		if err := c.BodyParser(newNewsTag); err != nil {
			log.Println(err)
			return err
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.ChangeNewsTag(id, newNewsTag, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была изменена", id)
		return c.SendString(message)

	})

	app.Delete("/newsTag/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.DeleteNewsTag(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была удалена", id)
		return c.SendString(message)
	})
	// ------- /newsTag -------------------------------------------
	//
	//
	//
	// ------- /news --------------------------------------------
	app.Get("/news", func(c *fiber.Ctx) error {
		message, err := workdb.SelectAllNews(conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(message)

	})

	app.Post("/news", func(c *fiber.Ctx) error {
		newNews := new(model.News)
		if err := c.BodyParser(newNews); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNews(newNews, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString("Новая запись была добавлена")
	})

	app.Get("/news/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		tag, err := workdb.SelectAnNews(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.SendString(tag)
	})

	app.Put("/news/:id", func(c *fiber.Ctx) error {
		newNews := new(model.News)
		if err := c.BodyParser(newNews); err != nil {
			log.Println(err)
			return err
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.ChangeNews(id, newNews, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была изменена", id)
		return c.SendString(message)

	})

	app.Delete("/news/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			log.Println(err)
			return err
		}
		err = workdb.DeleteNews(id, conn)
		if err != nil {
			log.Println(err)
			return err
		}
		message := fmt.Sprintf("Запись с id = %d была удалена", id)
		return c.SendString(message)
	})
	// ------- /news --------------------------------------------

	app.Get("/addNews", func(c *fiber.Ctx) error {

		return c.SendString("message")

	})

	app.Listen(":3000")

}
