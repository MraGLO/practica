package http

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetCategories(c *fiber.Ctx) error {
	categories, err := h.services.GetAllCategories()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	return c.JSON(categories)
}

func (h *Handlers) GetCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("categoryID"))
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

	category, err := h.services.GetCategoryByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if category.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(category)
}

func (h *Handlers) AddCategory(c *fiber.Ctx) error {
	var category model.Category

	err := json.Unmarshal(c.Body(), &category)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateCategoryData(category)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddCategory(category)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("categoryID")
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

	var category model.Category

	err = json.Unmarshal(c.Body(), &category)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	str := isValidateCategoryData(category)
	if str != "" {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateCategory(id, category)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}

func isValidateCategoryData(category model.Category) (str string) {
	str, b := isValidString(category.CategoryName)
	if category.CategoryName == "" || b {
		str = fmt.Sprintf("в русском названии: %s", str)
		log.Println(str)
		return
	}

	str, b = isValidString(category.CategoryNameEN)
	if category.CategoryNameEN == "" || b {
		str = fmt.Sprintf("в английском названии: %s", str)
		log.Println(str)
		return
	}
	return
}

func (h *Handlers) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("categoryID")
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

	found, err := h.services.DeleteCategory(id)
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
