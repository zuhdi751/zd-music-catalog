package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zuhdi751/zd_music_catalog/internal/configs"
	membershipsHandler "github.com/zuhdi751/zd_music_catalog/internal/handler/memberships"
	"github.com/zuhdi751/zd_music_catalog/internal/models/memberships"
	membershipsRepo "github.com/zuhdi751/zd_music_catalog/internal/repository/memberships"
	membershipsSvc "github.com/zuhdi751/zd_music_catalog/internal/service/memberships"
	"github.com/zuhdi751/zd_music_catalog/pkg/internalsql"
)

func main() {
	fmt.Println("Hello, world!")

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/", // for local configs file path
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatalf("failed to initialize configs: %v", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipRepo := membershipsRepo.NewRepository(db)
	membershipSvc := membershipsSvc.NewService(cfg, membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r, membershipSvc)
	membershipHandler.ResgisterRoute()

	r.Run(cfg.Service.Port)
}
