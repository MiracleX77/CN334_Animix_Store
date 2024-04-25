package main

import (
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/database"
	"github.com/MiracleX77/CN334_Animix_Store/user/entities"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	userMigrate(db)
}

func userMigrate(db database.Database) {
	db.GetDb().AutoMigrate(&entities.User{})
}
