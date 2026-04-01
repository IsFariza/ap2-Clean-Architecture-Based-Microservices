package repository

import (
	"context"

	"github.com/IsFariza/doctor-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type doctorRepository struct {
	collection *mongo.Collection
}

func NewDoctorRepository(client *mongo.Client) *doctorRepository {
	return &doctorRepository{
		collection: client.Database("doctor_db").Collection("doctors"),
	}
}

func (r *doctorRepository) Create(ctx context.Context, doctor *domain.Doctor) error {
	doc := FromDomain(doctor)

	res, err := r.collection.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		doctor.ID = objId.Hex()
	}
	return nil
}

func (r *doctorRepository) GetById(ctx context.Context, id string) (*domain.Doctor, error) {
	var doc DoctorDoc
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}

	if err := r.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		return nil, err
	}

	return doc.ToDomain(), nil
}

func (r *doctorRepository) GetAll(ctx context.Context) ([]*domain.Doctor, error) {
	var docs []DoctorDoc

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &docs); err != nil {
		return nil, err
	}

	result := make([]*domain.Doctor, len(docs))
	for i, doc := range docs {
		result[i] = doc.ToDomain()
	}

	return result, nil
}

func (r *doctorRepository) GetByEmail(ctx context.Context, email string) (*domain.Doctor, error) {
	var doc DoctorDoc
	filter := primitive.M{"email": email}
	if err := r.collection.FindOne(ctx, filter).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return doc.ToDomain(), nil
}
