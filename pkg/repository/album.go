package repository

import (
	"context"

	"github.com/Phaseant/MusicAPI/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AlbumRepo struct {
	db *mongo.Client
}

func NewAlbumRepo(db *mongo.Client) *AlbumRepo {
	return &AlbumRepo{db: db}
}

func (r *AlbumRepo) AddAlbum(album entity.Album) (string, error) {
	collection := r.db.Database(DBName).Collection(AlbumCol)
	album.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(context.Background(), album)

	if err != nil {
		return "", err
	}
	log.Info("repo.AddAlbum: added a new album with id: ", album.Id.Hex())
	return album.Id.Hex(), nil
}

func (r *AlbumRepo) GetAlbum(id string) (entity.Album, error) {
	collection := r.db.Database(DBName).Collection(AlbumCol)

	var album entity.Album

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Album{}, err
	}

	filter := bson.D{{Key: "_id", Value: objId}}
	err = collection.FindOne(context.Background(), filter).Decode(&album)

	if err != nil {
		return entity.Album{}, err
	}
	log.Info("repo.GetAlbum: returned album with id: ", album.Id.Hex())
	return album, nil
}

func (r *AlbumRepo) GetAllAlbums() ([]entity.Album, error) {
	collection := r.db.Database(DBName).Collection(AlbumCol)
	var albums []entity.Album

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return []entity.Album{}, err
	}

	for cursor.Next(context.Background()) {
		var album entity.Album
		err := cursor.Decode(&album)
		if err != nil {
			return []entity.Album{}, err
		}

		albums = append(albums, album)
	}

	if err := cursor.Err(); err != nil {
		return []entity.Album{}, err
	}
	log.Info("repo.GetAllAlbums: returned all albums from database")
	return albums, nil
}

func (r *AlbumRepo) DeleteAlbum(id string) bool {
	collection := r.db.Database(DBName).Collection(AlbumCol)

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err == nil
	}
	filter := bson.D{{Key: "_id", Value: objId}}

	result, _ := collection.DeleteOne(context.Background(), filter)

	if result.DeletedCount > 0 {
		log.Info("repo.DeleteAlbum: deleted album with id: ", id)
	}

	return result.DeletedCount != 0 //if deleted returns true
}
