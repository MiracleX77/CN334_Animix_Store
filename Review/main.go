package main

import (
	"github.com/MiracleX77/CN334_Animix_Store/configs"
	"github.com/MiracleX77/CN334_Animix_Store/database"
	"github.com/MiracleX77/CN334_Animix_Store/server"
)

func main() {
	cfg := configs.GetConfig()
	db := database.NewPostgresDatabase(&cfg)
	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
