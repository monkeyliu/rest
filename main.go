package main

import (
	"github.com/gin-gonic/gin"
	"github.com/monkeyliu/rest/service"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1/userinfo")
	{
		v1.GET("/", service.GETuser)
		v1.DELETE("/", service.DELETEuser)
		v1.PUT("/", service.UPDATEuser)
		v1.POST("/", service.CREATEuser)
	}
	r.Run()
}
