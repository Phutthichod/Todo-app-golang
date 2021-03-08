package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Label      string             `json:"label" bson:"label"`
	Check      bool               `json:"check" bson:"check"`
	UpdateTime time.Time          `json:"update_time" bson:"update_time"`
}
