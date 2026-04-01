package usecase

import (
	"context"
	"errors"

	"github.com/IsFariza/doctor-service/internal/domain"
)

type DoctorUsecase struct {
	doctorRepo domain.DoctorRepo
}

func NewDoctorUsecase(doctorRepo domain.DoctorRepo) *DoctorUsecase {
	return &DoctorUsecase{
		doctorRepo: doctorRepo,
	}
}

func (uc *DoctorUsecase) Create(ctx context.Context, doctor *domain.Doctor) error {
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

func (uc *DoctorUsecase) GetAll(ctx context.Context) ([]*domain.Doctor, error) {
	return uc.doctorRepo.GetAll(ctx)
}

func (uc *DoctorUsecase) GetByEmail(ctx context.Context, email string) (*domain.Doctor, error) {
	return uc.doctorRepo.GetByEmail(ctx, email)
}

func (uc *DoctorUsecase) GetById(ctx context.Context, id string) (*domain.Doctor, error) {
	return uc.doctorRepo.GetById(ctx, id)
}
