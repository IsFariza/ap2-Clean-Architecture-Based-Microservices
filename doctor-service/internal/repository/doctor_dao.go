package repository

import (
	"time"

	"github.com/IsFariza/doctor-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorDoc struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FullName       string             `bson:"full_name"`
	Specialization string             `bson:"specialization"`
	Email          string             `bson:"email"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}

func FromDoctorDomain(d *domain.Doctor) *DoctorDoc {
	doc := &DoctorDoc{
		FullName:       d.FullName,
		Specialization: d.Specialization,
		Email:          d.Email,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	if d.ID != "" {
		if objID, err := primitive.ObjectIDFromHex(d.ID); err == nil {
			doc.ID = objID
		}
	}
	return doc
}

func (d DoctorDoc) ToDomain() *domain.Doctor {
	return &domain.Doctor{
		ID:             d.ID.Hex(),
		FullName:       d.FullName,
		Specialization: d.Specialization,
		Email:          d.Email,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
	}
}
