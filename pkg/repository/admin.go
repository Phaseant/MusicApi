package repository

import (
	"context"

	"github.com/Phaseant/MusicAPI/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepo struct {
	db *mongo.Client
}

func NewAdminRepo(db *mongo.Client) *AdminRepo {
	return &AdminRepo{db: db}
}

func (r *AdminRepo) IsAdmin(id string) bool {
	collection := r.db.Database(DBName).Collection(AdminCol)
	objID, err := primitive.ObjectIDFromHex(id)

	log.Info(objID)
	if err != nil {
		return false
	}

	filter := bson.D{{Key: "_id", Value: objID}}

	counter, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Error(err)
		return false
	}

	if counter == 0 {
		return false
	}

	return true
}

func (r *AdminRepo) AddAdmin(admin entity.Admin) error {
	collection := r.db.Database(DBName).Collection(AdminCol)

	objId, err := primitive.ObjectIDFromHex(admin.Id)
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.TODO(), bson.D{{Key: "_id", Value: objId}})
	if err != nil {
		return err
	}
	return nil
}
