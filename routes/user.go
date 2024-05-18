package routes

import (
	controller "ShareShelf/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("api/v1/user/register", controller.RegisterUser)
	incomingRoutes.POST("api/v1/user/login", controller.LoginUser)

}
