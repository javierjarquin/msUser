package http

import (
	"msUser/internal/domain"
	usercase "msUser/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler maneja las solicitudes HTTP relacionadas con usuarios
type TandaHandler struct {
	TandaUC *usercase.TandaCase
}

// NewUserHandler crea una nueva instancia del controlador de usuarios
func NewTandaHandler(tandaUC *usercase.TandaCase) *TandaHandler {
	return &TandaHandler{TandaUC: tandaUC}
}

// CreateUser maneja la solicitud para crear un usuario
func (h *TandaHandler) CreateTanda(c *gin.Context) {
	var tanda domain.Tanda
	if err := c.ShouldBindJSON(&tanda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTanda, err := h.TandaUC.NewTanda(tanda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tanda"})
		return
	}

	c.JSON(http.StatusCreated, createdTanda)
}

func UpdateTandaHandler(tandaUC *usercase.TandaCase) *TandaHandler {
	return &TandaHandler{TandaUC: tandaUC}
}

func (h *TandaHandler) UpdateTanda(c *gin.Context) {
	var tanda domain.Tanda
	if err := c.ShouldBindJSON(&tanda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTanda, err := h.TandaUC.UpdateTanda(tanda)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	c.JSON(http.StatusOK, updatedTanda)
}

func GetTandaByIdHandler(tandaUC *usercase.TandaCase) *TandaHandler {
	return &TandaHandler{TandaUC: tandaUC}
}

func (h *TandaHandler) GetTandaById(c *gin.Context) {

	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	tanda, err := h.TandaUC.GetTandaById(idint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tanda no encontrada"})
		return
	}

	c.JSON(http.StatusOK, tanda)
}

func GetTandaByUserIdHandler(tandaUC *usercase.TandaCase) *TandaHandler {
	return &TandaHandler{TandaUC: tandaUC}
}

func (h *TandaHandler) GetTandaByUserId(c *gin.Context) {
	id := c.Param("id")
	idint, err := strconv.Atoi(id)
	tandas, err := h.TandaUC.GetTandaByUserId(idint)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, tandas)
}
