package router

import (
	"msUser/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *http.UserHandler) *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.POST("/update", userHandler.UpdateUser)
		userRoutes.GET("/:id", userHandler.GetUserById)
		userRoutes.POST("/login", userHandler.LoginUser)
	}
	return r
}
