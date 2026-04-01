package usecase

import (
	"context"
	"errors"

	"github.com/IsFariza/doctor-service/internal/domain"
)

type doctorUsecase struct {
	doctorRepo domain.DoctorRepo
}

func NewDoctorUsecase(doctorRepo domain.DoctorRepo) *doctorUsecase {
	return &doctorUsecase{
		doctorRepo: doctorRepo,
	}
}

func (uc *doctorUsecase) Create(ctx context.Context, doctor *domain.Doctor) error {
	if doctor.FullName == "" {
		return errors.New("full name is required")
	}
	if doctor.Email == "" {
		return errors.New("email is required")
	}

	exists, _ := uc.doctorRepo.GetByEmail(ctx, doctor.Email)
	if exists != nil {
		return errors.New("email must be unique")
	}
	return uc.doctorRepo.Create(ctx, doctor)
}

func (uc *doctorUsecase) GetAll(ctx context.Context) ([]*domain.Doctor, error) {
	return uc.doctorRepo.GetAll(ctx)
}

func (uc *doctorUsecase) GetByEmail(ctx context.Context, email string) (*domain.Doctor, error) {
	return uc.doctorRepo.GetByEmail(ctx, email)
}

func (uc *doctorUsecase) GetById(ctx context.Context, id string) (*domain.Doctor, error) {
	return uc.doctorRepo.GetById(ctx, id)
}
