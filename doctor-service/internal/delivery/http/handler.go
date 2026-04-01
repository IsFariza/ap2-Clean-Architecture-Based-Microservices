package http

import (
	"net/http"

	"github.com/IsFariza/doctor-service/internal/delivery/http/dto"
	"github.com/IsFariza/doctor-service/internal/domain"
	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	usecase domain.DoctorUsecase
}

func NewDoctorHandler(usecase domain.DoctorUsecase) *DoctorHandler {
	return &DoctorHandler{
		usecase: usecase,
	}
}

func (h *DoctorHandler) Create(c *gin.Context) {
	var req dto.DoctorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doctor := dto.ToDomain(req)
	if err := h.usecase.Create(c.Request.Context(), doctor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.FromDomain(doctor))
}

func (h *DoctorHandler) GetAll(c *gin.Context) {
	doctors, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	res := make([]dto.DoctorResponse, len(doctors))
	for i, d := range doctors {
		res[i] = dto.FromDomain(d)
	}
	c.JSON(http.StatusOK, res)

}

func (h *DoctorHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	doctor, err := h.usecase.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.FromDomain(doctor))
}
