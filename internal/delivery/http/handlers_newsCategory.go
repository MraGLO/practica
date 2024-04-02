package http

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetNewsCategories(c *fiber.Ctx) error {
	newsCategories, err := h.services.GetAllNewsCategories()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	return c.JSON(newsCategories)
}

func (h *Handlers) GetNewsCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsCategoryID"))
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}

	if id <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или равно 0"})
	}

	newsCategory, err := h.services.GetNewsCategoryByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if newsCategory.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(newsCategory)
}

func (h *Handlers) AddNewsCategory(c *fiber.Ctx) error {
	var newsCategory model.NewsCategory

	err := json.Unmarshal(c.Body(), &newsCategory)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateNewsCategoryData(newsCategory)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddNewsCategory(newsCategory)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateNewsCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsCategoryID")
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}

	if id <= 0 {
		log.Println("id <=0")
		c.SendStatus(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или рано 0"})
	}

	var newsCategory model.NewsCategory

	err = json.Unmarshal(c.Body(), &newsCategory)
	if err != nil {
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateNewsCategoryData(newsCategory)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateNewsCategory(id, newsCategory)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func isValidateNewsCategoryData(newsCategory model.NewsCategory) (str string) {
	if newsCategory.CategoryID <= 0 {
		str = "значение id категории невозможно"
		log.Println(str)
		return
	}

	if newsCategory.NewsID <= 0 {
		str = "значение id новости невозможно"
		log.Println(str)
		return
	}
	return
}

func (h *Handlers) DeleteNewsCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsCategoryId")
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}

	if id <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или равно 0"})
	}

	found, err := h.services.DeleteNewsCategory(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if !found {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	c.Status(200)
	return c.JSON(model.Error{Data: "Успешно удалено"})
}
