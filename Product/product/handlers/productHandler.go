package handlers

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	UpdateProduct(c echo.Context) error
	GetProductById(c echo.Context) error
	GetProductAll(c echo.Context) error
	DeleteProduct(c echo.Context) error
	GetProductAllByCategory(c echo.Context) error
	GetProductAllByName(c echo.Context) error
	InsertProduct(c echo.Context) error
}

type AuthorHandler interface {
	InsertAuthor(c echo.Context) error
	UpdateAuthor(c echo.Context) error
	GetAuthorById(c echo.Context) error
	GetAuthorAll(c echo.Context) error
	DeleteAuthor(c echo.Context) error
}

type PublisherHandler interface {
	InsertPublisher(c echo.Context) error
	UpdatePublisher(c echo.Context) error
	GetPublisherById(c echo.Context) error
	GetPublisherAll(c echo.Context) error
	DeletePublisher(c echo.Context) error
}

type CategoryHandler interface {
	InsertCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	GetCategoryById(c echo.Context) error
	GetCategoryAll(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type FavoriteHandler interface {
	InsertFavorite(c echo.Context) error
	GetFavoriteAllByUserId(c echo.Context) error
	DeleteFavorite(c echo.Context) error
}
