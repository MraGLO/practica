package http

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/MraGLO/practica/pkg/model"
	"github.com/gofiber/fiber/v2"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func (h *Handlers) GetNews(c *fiber.Ctx) error {
	news, err := h.services.GetAllNews()
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.JSON(news)
}

func (h *Handlers) GetNewsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsID"))
	if err != nil || id <= 0 {
		log.Println(err)
		return c.SendStatus(400)
	}

	news, err := h.services.GetNewsByID(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if news.ID == 0 {
		log.Println(err)
		return c.SendStatus(404)
	}

	return c.JSON(news)
}

func (h *Handlers) AddNews(c *fiber.Ctx) error {
	var news model.NewNews

	err := json.Unmarshal(c.Body(), &news)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}
	isLenCategories, isLenTags, err := isValidateNewNewsData(news)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	err = h.services.AddNews(news, isLenCategories, isLenTags)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно добавлено")
}

func (h *Handlers) UpdateNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsID")
	if err != nil || id <= 0 {
		log.Println(err)
		return c.SendStatus(400)
	}

	var news model.News

	err = json.Unmarshal(c.Body(), &news)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	if isValidateNewsData(news) != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	err = h.services.UpdateNews(id, news)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	c.Status(201)
	return c.SendString("Успешно обновлено")
}
func isValidateNewNewsData(news model.NewNews) (isLenCategories bool, isLenTags bool, err error) {
	str, b := isValidString(news.Shortname)
	if news.Shortname == "" || b {
		str = fmt.Sprintf("в shortname: %s", str)
		err = fmt.Errorf(str)
		log.Println(err)
		return
	}

	if news.Body == "" {
		err = fmt.Errorf("body пуст")
		log.Println(err)
		return
	}

	if news.PublishedTime == "" {
		err = fmt.Errorf("publishedTime пуст")
		log.Println(err)
		return
	}

	if len(news.Categories) != 0 {
		isLenCategories = true
	}

	if len(news.Tags) != 0 {
		isLenTags = true
	}

	return
}

func isValidateNewsData(news model.News) (err error) {
	str, b := isValidString(news.Shortname)
	if news.Shortname == "" || b {
		str = fmt.Sprintf("в shortname: %s", str)
		err = fmt.Errorf(str)
		log.Println(err)
		return
	}

	if news.Body == "" {
		err = fmt.Errorf("body пуст")
		log.Println(err)
		return
	}

	return
}

func (h *Handlers) DeleteNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsID")
	if err != nil || id <= 0 {
		log.Println(err)
		return c.SendStatus(400)
	}

	found, err := h.services.DeleteNews(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if !found {
		log.Println(err)
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func (h *Handlers) MdToHTMLNews(c *fiber.Ctx) error {
	var md []byte

	err := json.Unmarshal(c.Body(), &md)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	HTML, err := mdToHTML(md)
	if err != nil {
		log.Println(err)
		return c.SendStatus(400)
	}

	return c.JSON(HTML)

}

func mdToHTML(md []byte) (HTML []byte, err error) {
	var printAst = false
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	if printAst {
		fmt.Print("--- AST tree:\n")
		ast.Print(os.Stdout, doc)
		fmt.Print("\n")
	}

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	HTML = markdown.Render(doc, renderer)

	return
}
