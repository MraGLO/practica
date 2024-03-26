package http

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
	// "github.com/MraGLO/practica/pkg/model"
)

func (h *Handlers) GetCategories(c *fiber.Ctx) error {
	categories, err := h.services.GetAllCategories()
	if err != nil {
		return c.SendStatus(500)
	}
	return c.JSON(categories)
}

func (h *Handlers) GetCategoryById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	category, err := h.services.GetCategoryByID(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if category.ID == 0 {
		return c.SendStatus(404)
	}

	return c.JSON(category)
}

func (h *Handlers) AddCategory(c *fiber.Ctx) error {
	var category model.Category

	err := json.Unmarshal(c.Body(), &category)
	if err != nil {
		return c.SendStatus(400)
	}

	if isValidateCategoryData(category) != nil {
		return c.SendStatus(400)
	}

	err = h.services.AddCategory(category)
	if err != nil {
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно добавлено")
}

func (h *Handlers) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pointId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	var category model.Category

	err = json.Unmarshal(c.Body(), &category)
	if err != nil {
		return c.SendStatus(400)
	}

	if isValidateCategoryData(category) != nil {
		return c.SendStatus(400)
	}

	err = h.services.UpdateCategory(id, category)
	if err != nil {
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно обновлено")
}

func isValidateCategoryData(category model.Category) (err error) {
	str, b := isValidString(category.CategoryName)
	if category.CategoryName == "" || b {
		str = fmt.Sprintf("в русском названии: %s", str)
		err = fmt.Errorf(str)
		return
	}

	str, b = isValidString(category.CategoryNameEN)
	if category.CategoryNameEN == "" || b {
		str = fmt.Sprintf("в английском названии: %s", str)
		err = fmt.Errorf(str)
		return
	}
	return
}

func isValidString(s string) (str string, b bool) {
	b = false
	for _, char := range s {
		if s[0] == ' ' {
			b = true
			str = "первый символ является пробелом"
			return
		}
		if char == '\t' {
			b = true
			str = "в имени категррии приссутсвует табуляция"
			return
		}
		if char == '\n' {
			b = true
			str = "в имени категррии приссутсвует несколько строк"
			return
		}
	}
	return
}

func (h *Handlers) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("pointId")
	if err != nil || id <= 0 {
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteCategory(id)
	if err != nil {
		return c.SendStatus(500)
	}

	if !found {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
