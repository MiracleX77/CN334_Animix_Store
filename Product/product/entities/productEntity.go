package entities

import (
	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		AuthorId    int `json:"author_id"`
		Author      Author
		CategoryId  int `json:"category_id"`
		Category    Category
		PublisherId int `json:"publisher_id"`
		Publisher   Publisher
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		ImgUrl      string  `json:"img_url"`
		Status      string  `json:"status"`
	}

	UpdateProduct struct {
		gorm.Model
		AuthorId    int     `json:"author_id"`
		CategoryId  int     `json:"category_id"`
		PublisherId int     `json:"publisher_id"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		ImgUrl      string  `json:"img_url"`
		Status      string  `json:"status"`
	}

	InsertProduct struct {
		AuthorId    int     `json:"author_id"`
		CategoryId  int     `json:"category_id"`
		PublisherId int     `json:"publisher_id"`
		Name        string  `json:"name"`
		Description *string `json:"description"`
		Price       float64 `json:"price"`
		Stock       int     `json:"stock"`
		ImgUrl      string  `json:"img_url"`
		Status      string  `json:"status"`
	}
)
