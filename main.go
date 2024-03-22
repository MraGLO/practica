package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/MraGLO/practica/workdb"
)

func main() {

	dbpool := workdb.ConnectDB()

	defer dbpool.Close()

	app := fiber.New()
	app.Static("/", "./PublicationPage")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("main")
	})

	// ------- /category --------------------------------------------

	app.Get("/category", func(c *fiber.Ctx) error {
		categories, err := workdb.SelectAllCategory(dbpool)

		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(categories)

	})
	app.Post("/category", func(c *fiber.Ctx) error {
		newCategory := new(model.Category)
		if err := c.BodyParser(newCategory); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertCategory(newCategory, dbpool)
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
		category, err := workdb.SelectAnCategory(id, dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(category)
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
		err = workdb.ChangeCategory(id, newCategory, dbpool)
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
		err = workdb.DeleteCategory(id, dbpool)
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
		tags, err := workdb.SelectAllTags(dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(tags)

	})

	app.Post("/tag", func(c *fiber.Ctx) error {
		newTag := new(model.Tag)
		if err := c.BodyParser(newTag); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertTags(newTag, dbpool)
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
		tag, err := workdb.SelectAnTags(id, dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(tag)
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
		err = workdb.ChangeTags(id, newTag, dbpool)
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
		err = workdb.DeleteTags(id, dbpool)
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
		newsCategories, err := workdb.SelectAllNewsCategory(dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(newsCategories)

	})

	app.Post("/newsCategory", func(c *fiber.Ctx) error {
		newNewsCategory := new(model.NewsCategory)
		if err := c.BodyParser(newNewsCategory); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNewsCategory(newNewsCategory, dbpool)
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
		newCategory, err := workdb.SelectAnNewsCategory(id, dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(newCategory)
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
		err = workdb.ChangeNewsCategory(id, newNewsCategory, dbpool)
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
		err = workdb.DeleteNewsCategory(id, dbpool)
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
		newsTags, err := workdb.SelectAllNewsTag(dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(newsTags)

	})

	app.Post("/newsTag", func(c *fiber.Ctx) error {
		newNewsTag := new(model.NewsTag)
		if err := c.BodyParser(newNewsTag); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNewsTag(newNewsTag, dbpool)
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
		newTag, err := workdb.SelectAnNewsTag(id, dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(newTag)
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
		err = workdb.ChangeNewsTag(id, newNewsTag, dbpool)
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
		err = workdb.DeleteNewsTag(id, dbpool)
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
		news, err := workdb.SelectAllNews(dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(news)

	})

	app.Post("/news", func(c *fiber.Ctx) error {
		news := new(model.NewNews)
		if err := c.BodyParser(news); err != nil {
			log.Println(err)
			return err
		}
		err := workdb.InsertNews(news, dbpool)
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
		news, err := workdb.SelectAnNews(id, dbpool)
		if err != nil {
			log.Println(err)
			return err
		}
		return c.JSON(news)
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
		err = workdb.ChangeNews(id, newNews, dbpool)
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
		err = workdb.DeleteNews(id, dbpool)
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
