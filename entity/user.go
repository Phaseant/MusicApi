package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Username string             `json:"username" bson:"username" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}
