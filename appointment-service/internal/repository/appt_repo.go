package repository

import (
	"appointment-service/internal/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type appointmentRepo struct {
	collection *mongo.Collection
}

func NewAppointmentRepository(client *mongo.Client) domain.AppointmentRepo {
	return &appointmentRepo{
		collection: client.Database("appointment_db").Collection("appointments"),
	}
}

func (r *appointmentRepo) Create(ctx context.Context, appointment *domain.Appointment) error {
	doc := FromDomain(appointment)
	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if objId, ok := res.InsertedID.(bson.ObjectID); ok {
		appointment.ID = objId.Hex()
	}
	return nil
}

func (r *appointmentRepo) GetById(ctx context.Context, id string) (*domain.Appointment, error) {
	var doc AppointmentDoc
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}

	if err := r.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}
	return doc.ToDomain(), nil
}

func (r *appointmentRepo) GetAll(ctx context.Context) ([]*domain.Appointment, error) {
	var docs []AppointmentDoc

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := make([]*domain.Appointment, len(docs))
	for i, doc := range docs {
		result[i] = doc.ToDomain()
	}

	return result, nil
}

func (r *appointmentRepo) Update(ctx context.Context, id string, newStatus domain.Status) error {
	objID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	doc := FromDomain(&domain.Appointment{ID: id, Status: newStatus})
	update := bson.M{"$set": bson.M{
		"status":     doc.Status,
		"updated_at": doc.UpdatedAt,
	}}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("appointment not found")
	}
	return nil
}
