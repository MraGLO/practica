package app

import (
	"github.com/MraGLO/practica/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Get("category/all", handlers.GetCategories)
	router.Get("category/:categoryID", handlers.GetCategoryById)

	router.Get("tag/all", handlers.GetTags)
	router.Get("tag/:tagID", handlers.GetTagById)

	router.Get("newsCategory/all", handlers.GetNewsCategories)
	router.Get("newsCategory/:newsCategoryID", handlers.GetNewsCategoryById)

	router.Get("newsTag/all", handlers.GetNewsTags)
	router.Get("newsTag/:newsTagID", handlers.GetNewsTagById)

	router.Get("news/all", handlers.GetNews)
	router.Get("news/:newsID", handlers.GetNewsById)

	category := router.Group("/category")
	tag := router.Group("/tag")
	newsCategory := router.Group("/newsCategory")
	newsTag := router.Group("/newsTag")
	news := router.Group("/news")
	newNews := router.Group("/newNews")

	newNews.Post("add", handlers.AddNews)
	newNews.Post("mdToHTML", handlers.MdToHTMLNews)

	category.Post("/add", handlers.AddCategory)
	category.Put("/:categoryID", handlers.UpdateCategory)
	category.Delete("/:categoryID", handlers.DeleteCategory)

	tag.Post("/add", handlers.AddTag)
	tag.Put("/:tagID", handlers.UpdateTag)
	tag.Delete("/:tagID", handlers.DeleteTag)

	newsCategory.Post("/add", handlers.AddNewsCategory)
	newsCategory.Put("/:newsCategoryID", handlers.UpdateNewsCategory)
	newsCategory.Delete("/:newsCategoryID", handlers.DeleteNewsCategory)

	newsTag.Post("/add", handlers.AddNewsTag)
	newsTag.Put("/:newsTagID", handlers.UpdateNewsTag)
	newsTag.Delete("/:newsTagID", handlers.DeleteNewsTag)

	news.Post("/add", handlers.AddNews)
	news.Put("/:newsID", handlers.UpdateNews)
	news.Delete("/:newsID", handlers.DeleteNews)

}
