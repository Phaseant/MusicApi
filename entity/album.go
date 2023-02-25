package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"`
	Title       string             `bson:"title, omitempty"`
	Author      string             `bson:"author, omitempty"`
	Year        int                `bson:"year, omitempty"`
	CoverURL    string             `bson:"coverURL, omitempty"`
	Description string             `bson:"description, omitempty"`
	Duration    string             `bson:"duration, omitempty"`
	Songs       []Song             `bson:"songs, omitempty"`
	GeniousLink string             `bson:"geniousLink, omitempty"`
}

type Song struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Title    string             `bson:"title"`
	Duration string             `bson:"duration"`
}
