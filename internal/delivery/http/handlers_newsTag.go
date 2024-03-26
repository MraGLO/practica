package http

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetNewsTags(c *fiber.Ctx) error {
	newsTags, err := h.services.GetAllNewsTags()
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}
	return c.JSON(newsTags)
}

func (h *Handlers) GetNewsTagById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsTagID"))
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	newTag, err := h.services.GetNewsTagByID(id)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	if newTag.ID == 0 {
		log.Panicln(err)
		return c.SendStatus(404)
	}

	return c.JSON(newTag)
}

func (h *Handlers) AddNewsTag(c *fiber.Ctx) error {
	var newsTag model.NewsTag

	err := json.Unmarshal(c.Body(), &newsTag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateNewsTagData(newsTag) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.AddNewsTag(newsTag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно добавлено")
}

func (h *Handlers) UpdateNewsTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newTagID")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	var newsTag model.NewsTag

	err = json.Unmarshal(c.Body(), &newsTag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateNewsTagData(newsTag) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.UpdateNewsTag(id, newsTag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно обновлено")
}

func isValidateNewsTagData(newsTag model.NewsTag) (err error) {
	if newsTag.TagID <= 0 {
		err = fmt.Errorf("значение id тега невозможно")
		log.Panicln(err)
		return
	}

	if newsTag.NewsID <= 0 {
		err = fmt.Errorf("значение id новости невозможно")
		log.Panicln(err)
		return
	}
	return
}

func (h *Handlers) DeleteNewsTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsTagId")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteNewsTag(id)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	if !found {
		log.Panicln(err)
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
