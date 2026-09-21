package router

import (
	"github.com/cris329/aprendiz-postgres/internal/controller"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func New(ctrl *controller.AprendizController) *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())
	api := r.Group("/api/v1/aprendiz_postgres")
	{
		api.POST("", ctrl.Create)
		api.GET("", ctrl.GetAll)
		api.GET("/:id", ctrl.GetByID)
		api.PUT("/:id", ctrl.Update)
		api.DELETE("/:id", ctrl.Delete)
	}
	return r
}
