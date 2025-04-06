package router

import (
	"msUser/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *http.UserHandler
}

func NewRouter(handler *http.UserHandler) *Router {
	return &Router{handler: handler}
}

func (r *Router) RegisterRoutes(engine *gin.Engine) {

	userRoutes := engine.Group("/users")
	{
		userRoutes.POST("/", r.handler.CreateUser)
		userRoutes.POST("/update", r.handler.UpdateUser)
		userRoutes.GET("/:id", r.handler.GetUserById)
		userRoutes.POST("/login", r.handler.LoginUser)
	}

}

// registerRoutes defines the routes for the applicatio
