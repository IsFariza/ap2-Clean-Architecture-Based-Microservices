package domain

import "context"

type DoctorRepo interface {
	Create(ctx context.Context, doctor *Doctor) error
	GetById(ctx context.Context, id string) (*Doctor, error)
	GetAll(ctx context.Context) ([]*Doctor, error)
	GetByEmail(ctx context.Context, emial string) (*Doctor, error)
}
