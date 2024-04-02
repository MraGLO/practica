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
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}
	return c.JSON(news)
}

func (h *Handlers) GetNewsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("newsID"))
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

	news, err := h.services.GetNewsByID(id)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	if news.ID == 0 {
		log.Println(err)
		c.Status(404)
		return c.JSON(model.Error{Data: "Данных по данному id не существует"})
	}

	return c.JSON(news)
}

func (h *Handlers) AddNews(c *fiber.Ctx) error {
	var news model.NewNews

	err := json.Unmarshal(c.Body(), &news)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}
	isLenCategories, isLenTags, str := isValidateNewNewsData(news)
	if str != "" {
		log.Println(str)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.AddNews(news, isLenCategories, isLenTags)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно добавлено"})
}

func (h *Handlers) UpdateNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsID")
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

	var news model.News

	err = json.Unmarshal(c.Body(), &news)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}
	str := isValidateNewsData(news)
	if str != "" {
		log.Println(str)
		c.Status(400)
		return c.JSON(model.Error{Data: str})
	}

	err = h.services.UpdateNews(id, news)
	if err != nil {
		log.Println(err)
		c.Status(500)
		return c.JSON(model.Error{Data: "Невозможно обратиться к серверу"})
	}

	c.Status(201)
	return c.JSON(model.Error{Data: "Успешно обновлено"})
}
func isValidateNewNewsData(news model.NewNews) (isLenCategories bool, isLenTags bool, str string) {
	str, b := isValidString(news.Shortname)
	if news.Shortname == "" || b {
		str = fmt.Sprintf("в shortname: %s", str)
		log.Println(str)
		return
	}

	if news.Body == "" {
		str = "body пуст"
		log.Println(str)
		return
	}

	if news.PublishedTime == "" {
		str = "publishedTime пуст"
		log.Println(str)
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

func isValidateNewsData(news model.News) (str string) {
	str, b := isValidString(news.Shortname)
	if news.Shortname == "" || b {
		str = fmt.Sprintf("в shortname: %s", str)
		log.Println(str)
		return
	}

	if news.Body == "" {
		str = "body пуст"
		log.Println(str)
		return
	}

	return
}

func (h *Handlers) DeleteNews(c *fiber.Ctx) error {
	id, err := c.ParamsInt("newsID")
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

	found, err := h.services.DeleteNews(id)
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

func (h *Handlers) MdToHTMLNews(c *fiber.Ctx) error {
	var md model.MdAndHtml

	err := json.Unmarshal(c.Body(), &md)
	if err != nil {
		log.Println(err)
		c.Status(400)
		return c.JSON(model.Error{Data: "ошибка в отправленном json"})
	}

	HTMLstr := mdToHTML([]byte(md.Data))
	var HTML = model.MdAndHtml{Data: string(HTMLstr)}

	return c.JSON(HTML)

}

func mdToHTML(md []byte) (HTML []byte) {
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
