package http

import (
	"github.com/IsFariza/doctor-service/internal/repository"
	"github.com/IsFariza/doctor-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(mongoClient *mongo.Client) *gin.Engine {
	r := gin.Default()

	repo := repository.NewDoctorRepository(mongoClient)
	uc := usecase.NewDoctorUsecase(repo)
	handler := NewDoctorHandler(uc)

	r.POST("/doctors", handler.Create)
	r.GET("/doctors", handler.GetAll)
	r.GET("/doctors/:id", handler.GetById)

	return r
}
