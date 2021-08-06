package api

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Online",
	})
}

func start_api() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		events := v1.Group("/events")
		{
			events.GET("/ping", ping)
		}
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/ping", ping)
		}
		news := v1.Group("/news")
		{
			news.GET("/ping", ping)
		}
	}
	router.Run()
}