package main

import (
	"fmt"
	"log"
	"wxcloudrun-golang/internal/app/service"
	"wxcloudrun-golang/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	service := service.NewService()
	router := gin.Default()
	router.GET("/auth/login", service.WeChatLogin)
	router.GET("/courts", service.GetCounts)
	router.GET("/courts/:id", service.GetCountInfo)

	router.POST("/events", service.StartEvent)
	router.POST("/collects/:fileID", service.ToggleCollectVideo)
	router.GET("/user/collects", service.GetCollectVideos)

	router.GET("/user/events", service.GetEventVideos)
	router.GET("/recommend/videos", service.GetRecommendVideos)

	log.Fatal(router.Run())
}
