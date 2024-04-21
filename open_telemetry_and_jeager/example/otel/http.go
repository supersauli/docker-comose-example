package otel

import "github.com/gin-gonic/gin"

func Run() {
	Init("gin-otel")
	base := gin.Default()
	r := base.Use(Count)
	r.GET("/add", Add)

	base.Run(":8080")
}
