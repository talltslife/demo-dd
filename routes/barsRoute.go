package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/talltslife/demo-dd/controllers"
)

func InitilizeBarRoutes(r *gin.Engine) {
	barRoutes := r.Group("/bars")
	{
		barRoutes.GET("/search", controllers.SearchBars)
	}
}
