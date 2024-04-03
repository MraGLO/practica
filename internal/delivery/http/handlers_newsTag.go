package http

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetNewsTags(c *fiber.Ctx) error {
	newsTags, err := h.services.GetAllNewsTags()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	return c.JSON(newsTags)
}

func (h *Handlers) GetNewsTagById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsTagID"))
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

	newsTag, err := h.services.GetNewsTagByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if newsTag.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(newsTag)
}

func (h *Handlers) AddNewsTag(c *fiber.Ctx) error {
	var newsTag model.NewsTag

	err := json.Unmarshal(c.Body(), &newsTag)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateNewsTagData(newsTag)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddNewsTag(newsTag)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateNewsTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsTagID")
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "id должно быть числом больше 0"})
	}

	if id <= 0 {
		log.Println("id <=0")
		c.Status(400)
		return c.JSON(model.Error{Data: "id не может быть меньше или рано 0"})
	}

	var newsTag model.NewsTag

	err = json.Unmarshal(c.Body(), &newsTag)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateNewsTagData(newsTag)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateNewsTag(id, newsTag)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func isValidateNewsTagData(newsTag model.NewsTag) (str string) {
	if newsTag.TagID <= 0 {
		str = "значение id тега невозможно"
		log.Println(str)
		return
	}

	if newsTag.NewsID <= 0 {
		str = "значение id новости невозможно"
		log.Println(str)
		return
	}
	return
}

func (h *Handlers) DeleteNewsTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsTagId")
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

	found, err := h.services.DeleteNewsTag(id)
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
