package usecase

import (
	"appointment-service/internal/domain"
	"context"
	"errors"
)

type appointmentUsecase struct {
	repo         domain.AppointmentRepo
	doctorClient domain.DoctorClient
}

func NewAppointmentUsecase(repo domain.AppointmentRepo, dc domain.DoctorClient) domain.AppointmentUsecase {
	return &appointmentUsecase{
		repo:         repo,
		doctorClient: dc,
	}
}

func (uc *appointmentUsecase) Create(ctx context.Context, appt *domain.Appointment) error {
	if appt.Title == "" || appt.DoctorID == "" {
		return errors.New("title, doctor_id are required")
	}

	exists, error := uc.doctorClient.DoctorExists(ctx, appt.DoctorID)
	if error != nil {
		return error
	}
	if !exists {
		return errors.New("doctor does not exist")
	}

	appt.Status = domain.StatusNew
	return uc.repo.Create(ctx, appt)
}

func (uc *appointmentUsecase) GetById(ctx context.Context, id string) (*domain.Appointment, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *appointmentUsecase) GetAll(ctx context.Context) ([]*domain.Appointment, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *appointmentUsecase) Update(ctx context.Context, id string, newStatus domain.Status) error {
	currentStatus, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if currentStatus.Status == domain.StatusDone && newStatus == domain.StatusNew {
		return errors.New("cannot change status from done to new")
	}
	if newStatus != domain.StatusDone && newStatus != domain.StatusInProgress &&
		newStatus != domain.StatusNew {
		return errors.New("invalid status")
	}
	return uc.repo.Update(ctx, id, newStatus)
}
