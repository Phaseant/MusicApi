package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	Id          primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	Title       string             `bson:"title, omitempty" json:"title,omitempty"`
	Author      string             `bson:"author, omitempty" json:"author,omitempty"`
	Year        int                `bson:"year, omitempty" json:"year,omitempty"`
	CoverURL    string             `bson:"coverURL, omitempty" json:"coverURL,omitempty"`
	Description string             `bson:"description, omitempty" json:"description,omitempty"`
	Duration    string             `bson:"duration, omitempty" json:"duration,omitempty"`
	Songs       []Song             `bson:"songs, omitempty" json:"songs,omitempty"`
	GeniousLink string             `bson:"geniousLink, omitempty" json:"geniousLink,omitempty"`
}

type Song struct {
	// Id       primitive.ObjectID `bson:"_id, omitempty" json:"omitempty"`
	Title    string `bson:"title" json:"title,omitempty"`
	Duration string `bson:"duration" json:"duration,omitempty"`
}
