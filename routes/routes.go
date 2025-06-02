package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/anuprajwal/go-mod/controllers"
	// "github.com/anuprajwal/go-mod/comand"
	
)

func RegesteredRoutes(router *gin.Engine){
	api := router.Group("/api")
	api.GET("/ping",controllers.Ping)
	api.POST("/comand-management",controllers.CommandManaer)
}