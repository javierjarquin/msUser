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

type TandaRouter struct {
	handler *http.TandaHandler
}

func NewTandaRouter(handler *http.TandaHandler) *TandaRouter {
	return &TandaRouter{handler: handler}
}

func (r *TandaRouter) RegisterRoutes(engine *gin.Engine) {
	// Define las rutas para la entidad Tanda
	tandaRoutes := engine.Group("/tandas")
	{
		tandaRoutes.POST("/", r.handler.CreateTanda)
		tandaRoutes.POST("/update", r.handler.UpdateTanda)
		tandaRoutes.GET("/:id", r.handler.GetTandaById)
		tandaRoutes.GET("/user/:id", r.handler.GetTandaByUserId)
	}
}

// Agregamos las rutas para el tanda usuario
type TandaUsuarioRouter struct {
	handler *http.TandaUsuarioHandler
}

func NewTandaUsuarioRouter(handler *http.TandaUsuarioHandler) *TandaUsuarioRouter {
	return &TandaUsuarioRouter{handler: handler}
}
func (r *TandaUsuarioRouter) RegisterRoutes(engine *gin.Engine) {
	// Define las rutas para la entidad TandaUsuario
	tandaUsuarioRoutes := engine.Group("/tandaUsuarios")
	{
		tandaUsuarioRoutes.POST("/", r.handler.CreateTandaUsuario)
		tandaUsuarioRoutes.POST("/update", r.handler.UpdateTandaUsuario)
		tandaUsuarioRoutes.GET("/:id", r.handler.GetTandaUsuarioByTandaId)
	}
}

// Agregamos las rutas para el tanda pago
type TandaPagoRouter struct {
	handler *http.TandaPagoHandler
}

func NewTandaPagoRouter(handler *http.TandaPagoHandler) *TandaPagoRouter {
	return &TandaPagoRouter{handler: handler}
}
func (r *TandaPagoRouter) RegisterRoutes(engine *gin.Engine) {
	// Define las rutas para la entidad TandaPago
	tandaPagoRoutes := engine.Group("/tandaPagos")
	{
		tandaPagoRoutes.POST("/", r.handler.CreateTandaPago)
		tandaPagoRoutes.POST("/update", r.handler.UpdateTandaPago)
		tandaPagoRoutes.GET("/:id", r.handler.GetTandaPagoByTandaUsuarioId)
	}
}
