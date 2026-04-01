package domain

import "context"

type DoctorUsecase interface {
	Create(ctx context.Context, doctor *Doctor) error
	GetById(ctx context.Context, id string) (*Doctor, error)
	GetAll(ctx context.Context) ([]*Doctor, error)
}
