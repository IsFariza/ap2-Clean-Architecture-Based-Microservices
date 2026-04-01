package domain

import "context"

type AppointmentRepo interface {
	Create(ctx context.Context, appointment *Appointment) error
	GetById(ctx context.Context, id string) (*Appointment, error)
	GetAll(ctx context.Context) ([]*Appointment, error)
	Update(ctx context.Context, id string, newStatus Status) error
}
