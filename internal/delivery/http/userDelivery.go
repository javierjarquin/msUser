package http

import (
	"msUser/internal/domain"
	usercase "msUser/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler maneja las solicitudes HTTP relacionadas con usuarios
type UserHandler struct {
	UserUC *usercase.UserCase
}

// NewUserHandler crea una nueva instancia del controlador de usuarios
func NewUserHandler(userUC *usercase.UserCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

// CreateUser maneja la solicitud para crear un usuario
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := h.UserUC.NewUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el usuario"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func UpdateUserHandler(userUC *usercase.UserCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := h.UserUC.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func GetUserByIdHandler(userUC *usercase.UserCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

func (h *UserHandler) GetUserById(c *gin.Context) {

	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	user, err := h.UserUC.GetUserById(idint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func LoginUserHandler(userUC *usercase.UserCase) *UserHandler {
	return &UserHandler{UserUC: userUC}
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var login domain.User
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := h.UserUC.LoginUser(login.Email, login.Pass, c.ClientIP(), "ogin.Comments")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, session)
}
