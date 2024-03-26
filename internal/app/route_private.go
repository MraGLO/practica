package app

import (
	"github.com/MraGLO/practica/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(router *fiber.App, handlers *http.Handlers) {
	router.Get("category/all", handlers.GetCategories)
	router.Get("category/:categoryID", handlers.GetCategoryById)

	router.Get("tag/all", handlers.GetCategories)
	router.Get("tag/:tagID", handlers.GetCategoryById)

	router.Get("newsCategory/all", handlers.GetCategories)
	router.Get("newsCategory/:newsCategoryID", handlers.GetCategoryById)

	router.Get("newsTag/all", handlers.GetCategories)
	router.Get("newsTag/:newsTagID", handlers.GetCategoryById)

	router.Get("news/all", handlers.GetCategories)
	router.Get("news/:newsID", handlers.GetCategoryById)

	category := router.Group("/category")
	tag := router.Group("/tag")
	newsCategory := router.Group("/newsCategory")
	newsTag := router.Group("/newsTag")
	news := router.Group("/news")

	category.Post("/add", handlers.AddCategory)
	category.Put("/:categoryID", handlers.UpdateCategory)
	category.Delete("/:categoryID", handlers.DeleteCategory)

	tag.Post("/add", handlers.AddTag)
	tag.Put("/:putID", handlers.UpdateTag)
	tag.Delete("/:putID", handlers.DeleteTag)

	newsCategory.Post("/add", handlers.AddNewsCategory)
	newsCategory.Put("/:newsCategoryID", handlers.UpdateNewsCategory)
	newsCategory.Delete("/:newsCategoryID", handlers.DeleteNewsCategory)

	newsTag.Post("/add", handlers.AddNewsTag)
	newsTag.Put("/:newTagID", handlers.UpdateNewsTag)
	newsTag.Delete("/:newTagID", handlers.DeleteNewsTag)

	news.Post("/add", handlers.AddNews)
	news.Put("/:newsID", handlers.UpdateNews)
	news.Delete("/:newsID", handlers.DeleteNews)

}
