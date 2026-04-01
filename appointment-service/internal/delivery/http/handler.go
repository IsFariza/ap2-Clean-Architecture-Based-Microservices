package delivery

import (
	"appointment-service/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	usecase domain.AppointmentUsecase
}

func NewAppointmentHandler(usecase domain.AppointmentUsecase) *appointmentHandler {
	return &appointmentHandler{
		usecase: usecase,
	}
}

func (h *appointmentHandler) Create(c *gin.Context) {
	var req CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appt := ToDomain(req)
	if err := h.usecase.Create(c.Request.Context(), appt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, FromDomain(appt))
}

func (h *appointmentHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	appt, error := h.usecase.GetById(c.Request.Context(), id)
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
		return
	}
	c.JSON(http.StatusOK, FromDomain(appt))
}

func (h *appointmentHandler) GetAll(c *gin.Context) {
	appointments, err := h.usecase.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]AppointmentResponse, len(appointments))
	for i, appt := range appointments {
		res[i] = FromDomain(appt)
	}
	c.JSON(http.StatusOK, res)
}

func (h *appointmentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if error := c.ShouldBindJSON(&req); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status is required"})
		return
	}
	if error := h.usecase.Update(c.Request.Context(), id, domain.Status(req.Status)); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "appointment updated successfully"})
}
