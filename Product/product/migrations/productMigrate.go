package main

import (
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/database"
	"github.com/MiracleX77/CN334_Animix_Store/product/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	addressMigrate(db)
}

func addressMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.Author{})
	db.GetDb().AutoMigrate(&entities.Publisher{})
	db.GetDb().AutoMigrate(&entities.Category{})
	db.GetDb().AutoMigrate(&entities.Product{})
	db.GetDb().AutoMigrate(&entities.Favorite{})
}
