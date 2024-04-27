package models

import "time"

type InsertProductModel struct {
	AuthorId    uint64  `json:"author_id" validate:"required"`
	CategoryId  uint64  `json:"category_id" validate:"required"`
	PublisherId uint64  `json:"publisher_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Img         string  `json:"img" `
}

type UpdateProductModel struct {
	ID          uint64  `json:"id" validate:"required" `
	AuthorId    uint64  `json:"author_id" validate:"required"`
	CategoryId  uint64  `json:"category_id" validate:"required"`
	PublisherId uint64  `json:"publisher_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Img         string  `json:"img" validate:"required"`
}

type ProductModel struct {
	ID          uint64         `json:"id"`
	Author      AuthorModel    `json:"author"`
	Category    CategoryModel  `json:"category"`
	Publisher   PublisherModel `json:"publisher"`
	Name        string         `json:"name"`
	Description *string        `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	ImgUrl      string         `json:"img_url"`
	Status      string         `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type AuthorModel struct {
	ID          uint64  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}
type InsertAuthorModel struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type CategoryModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type InsertCategoryModel struct {
	Name string `json:"name" validate:"required"`
}

type PublisherModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type InsertPublisherModel struct {
	Name string `json:"name" validate:"required"`
}

type FavoriteModel struct {
	ID        uint64 `json:"id"`
	Product   ProductFavoriteModel
	UserId    uint64 `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductFavoriteModel struct {
	ID     uint64  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	ImgUrl string  `json:"img_url"`
}

type InsertFavoriteModel struct {
	ProductId uint64 `json:"product_id" validate:"required"`
	UserId    uint64 `json:"user_id" validate:"required"`
	Status    string `json:"status"`
}
