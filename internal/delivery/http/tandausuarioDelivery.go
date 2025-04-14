package http

import (
	"msUser/internal/domain"
	usercase "msUser/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler maneja las solicitudes HTTP relacionadas con usuarios
type TandaUsuarioHandler struct {
	TandaUsuarioUC *usercase.TandaUsuarioCase
}

// NewUserHandler crea una nueva instancia del controlador de usuarios
func NewTandaUsuarioHandler(tandausuarioUC *usercase.TandaUsuarioCase) *TandaUsuarioHandler {
	return &TandaUsuarioHandler{TandaUsuarioUC: tandausuarioUC}
}

// CreateUser maneja la solicitud para crear un usuario
func (h *TandaUsuarioHandler) CreateTandaUsuario(c *gin.Context) {
	var tandausuario domain.TandaUsuario
	if err := c.ShouldBindJSON(&tandausuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTandaUsuario, err := h.TandaUsuarioUC.NewTandaUsuario(tandausuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tanda"})
		return
	}

	c.JSON(http.StatusCreated, createdTandaUsuario)
}

func UpdateTandaUsuarioHandler(tandausuarioUC *usercase.TandaUsuarioCase) *TandaUsuarioHandler {
	return &TandaUsuarioHandler{TandaUsuarioUC: tandausuarioUC}
}

func (h *TandaUsuarioHandler) UpdateTandaUsuario(c *gin.Context) {
	var tandausuario domain.TandaUsuario
	if err := c.ShouldBindJSON(&tandausuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTandaUsuario, err := h.TandaUsuarioUC.UpdateTandaUsuario(tandausuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, updatedTandaUsuario)
}

func GetTandaUsuarioByTandaIdHandler(tandausuarioUC *usercase.TandaUsuarioCase) *TandaUsuarioHandler {
	return &TandaUsuarioHandler{TandaUsuarioUC: tandausuarioUC}
}

func (h *TandaUsuarioHandler) GetTandaUsuarioByTandaId(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	tandausuario, err := h.TandaUsuarioUC.GetTandaUsuarioByTandaId(idint)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, tandausuario)
}
