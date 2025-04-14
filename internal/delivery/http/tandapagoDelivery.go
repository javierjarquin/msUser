package http

import (
	"msUser/internal/domain"
	usercase "msUser/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler maneja las solicitudes HTTP relacionadas con usuarios
type TandaPagoHandler struct {
	TandaPagoUC *usercase.TandaPagoCase
}

// NewUserHandler crea una nueva instancia del controlador de usuarios
func NewTandaPagoHandler(tandapagoUC *usercase.TandaPagoCase) *TandaPagoHandler {
	return &TandaPagoHandler{TandaPagoUC: tandapagoUC}
}

// CreateUser maneja la solicitud para crear un usuario
func (h *TandaPagoHandler) CreateTandaPago(c *gin.Context) {
	var tandapago domain.TandaPago
	if err := c.ShouldBindJSON(&tandapago); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTandaPago, err := h.TandaPagoUC.NewTandaPago(tandapago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tanda"})
		return
	}

	c.JSON(http.StatusCreated, createdTandaPago)
}

func UpdateTandaPagoHandler(tandapagoUC *usercase.TandaPagoCase) *TandaPagoHandler {
	return &TandaPagoHandler{TandaPagoUC: tandapagoUC}
}

func (h *TandaPagoHandler) UpdateTandaPago(c *gin.Context) {
	var tandapago domain.TandaPago
	if err := c.ShouldBindJSON(&tandapago); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTandaPago, err := h.TandaPagoUC.UpdateTandaPago(tandapago)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, updatedTandaPago)
}

func GetTandaPagoByTandaUsuarioIdHandler(tandapagoUC *usercase.TandaPagoCase) *TandaPagoHandler {
	return &TandaPagoHandler{TandaPagoUC: tandapagoUC}
}

func (h *TandaPagoHandler) GetTandaPagoByTandaUsuarioId(c *gin.Context) {
	id := c.Param("tandausuarioId")
	idint, err := strconv.Atoi(id)
	tandapago, err := h.TandaPagoUC.GetTandaPagoByTandaUsuarioId(idint)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, tandapago)
}
