package delivery

import (
	"appointment-service/internal/repository"
	"appointment-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(mongoClient *mongo.Client, doctorURL string) *gin.Engine {
	r := gin.Default()

	repo := repository.NewAppointmentRepository(mongoClient)
	doctorClient := NewDoctorClient(doctorURL)
	uc := usecase.NewAppointmentUsecase(repo, doctorClient)
	handler := NewAppointmentHandler(uc)

	r.POST("/appointments", handler.Create)
	r.GET("/appointments", handler.GetAll)
	r.GET("/appointments/:id", handler.GetById)
	r.PATCH("/appointments/:id/status", handler.Update)

	return r
}
