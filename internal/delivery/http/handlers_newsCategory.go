package http

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetNewsCategories(c *fiber.Ctx) error {
	newsCategories, err := h.services.GetAllNewsCategories()
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}
	return c.JSON(newsCategories)
}

func (h *Handlers) GetNewsCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsCategoryID"))
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	newCategory, err := h.services.GetNewsCategoryByID(id)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	if newCategory.ID == 0 {
		log.Panicln(err)
		return c.SendStatus(404)
	}

	return c.JSON(newCategory)
}

func (h *Handlers) AddNewsCategory(c *fiber.Ctx) error {
	var newCategory model.NewsCategory

	err := json.Unmarshal(c.Body(), &newCategory)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateNewsCategoryData(newCategory) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.AddNewsCategory(newCategory)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно добавлено")
}

func (h *Handlers) UpdateNewsCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newCategoryID")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	var newsCategory model.NewsCategory

	err = json.Unmarshal(c.Body(), &newsCategory)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateNewsCategoryData(newsCategory) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.UpdateNewsCategory(id, newsCategory)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно обновлено")
}

func isValidateNewsCategoryData(newsCategory model.NewsCategory) (err error) {
	if newsCategory.CategoryID <= 0 {
		err = fmt.Errorf("значение id категории невозможно")
		log.Panicln(err)
		return
	}

	if newsCategory.NewsID <= 0 {
		err = fmt.Errorf("значение id новости невозможно")
		log.Panicln(err)
		return
	}
	return
}

func (h *Handlers) DeleteNewsCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsCategoryId")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteNewsCategory(id)
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
