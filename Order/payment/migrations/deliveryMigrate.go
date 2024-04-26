package main

import (
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/database"
	"github.com/MiracleX77/CN334_Animix_Store/payment/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	addressMigrate(db)
}

func addressMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.Payment{})
}
