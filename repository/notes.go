package repository

import (
	"context"
	"time"
	"todo-go/config"
	"todo-go/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Noter interface {
	ListNote() ([]model.Note, error)
	AddNote(note model.Note) (interface{}, error)
	DeleteNote(id string) error
	UpdateNote(id string, note model.Note) error
}

type NoteRepository struct {
	ConnectionDB *mongo.Client
}

const collectionName = "note"

func (noteRepo *NoteRepository) AddNote(note model.Note) (interface{}, error) {
	note.UpdateTime = time.Now()
	id, err := noteRepo.ConnectionDB.Database(config.DatabaseName).Collection(collectionName).InsertOne(context.TODO(), note)
	return id, err
}

func (noteRepo *NoteRepository) ListNote() ([]model.Note, error) {
	var notes []model.Note
	cursor, err := noteRepo.ConnectionDB.Database(config.DatabaseName).Collection(collectionName).Find(context.TODO(), bson.D{})
	if err != nil {
		return notes, err
	}
	err = cursor.All(context.TODO(), &notes)
	return notes, err

}
func (noteRepo *NoteRepository) DeleteNote(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	_, err = noteRepo.ConnectionDB.Database(config.DatabaseName).Collection(collectionName).DeleteOne(context.TODO(), bson.M{"_id": objectId})
	return err
}
func (noteRepo *NoteRepository) UpdateNote(id string, note model.Note) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	note.UpdateTime = time.Now()
	update := bson.M{
		"$set": note,
	}
	_, err = noteRepo.ConnectionDB.Database(config.DatabaseName).Collection(collectionName).UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)
	return err
}
