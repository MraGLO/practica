package http

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetTags(c *fiber.Ctx) error {
	tags, err := h.services.GetAllTags()
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}
	return c.JSON(tags)
}

func (h *Handlers) GetTagById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("tagID"))
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	tag, err := h.services.GetTagByID(id)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	if tag.ID == 0 {
		log.Panicln(err)
		return c.SendStatus(404)
	}

	return c.JSON(tag)
}

func (h *Handlers) AddTag(c *fiber.Ctx) error {
	var tag model.Tag

	err := json.Unmarshal(c.Body(), &tag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateTagData(tag) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.AddTags(tag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно добавлено")
}

func (h *Handlers) UpdateTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("tagID")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	var tag model.Tag

	err = json.Unmarshal(c.Body(), &tag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	if isValidateTagData(tag) != nil {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	err = h.services.UpdateTag(id, tag)
	if err != nil {
		log.Panicln(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно обновлено")
}

func isValidateTagData(tag model.Tag) (err error) {
	str, b := isValidString(tag.TagName)
	if tag.TagName == "" || b {
		str = fmt.Sprintf("в русском названии: %s", str)
		err = fmt.Errorf(str)
		log.Panicln(err)
		return
	}

	str, b = isValidString(tag.TagNameEN)
	if tag.TagNameEN == "" || b {
		str = fmt.Sprintf("в английском названии: %s", str)
		err = fmt.Errorf(str)
		log.Panicln(err)
		return
	}
	return
}

func (h *Handlers) DeleteTag(c *fiber.Ctx) error {
	id, err := c.ParamsInt("tagID")
	if err != nil || id <= 0 {
		log.Panicln(err)
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteTag(id)
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
