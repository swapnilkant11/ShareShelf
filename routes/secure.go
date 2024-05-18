package routes

import (
	controller "ShareShelf/controllers"

	"github.com/gin-gonic/gin"
)

func SecureRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("api/v1/secure", controller.Ping)
}
