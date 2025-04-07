package model

import (
	//mongodb ids are bsons basically, common practice, fetch the primitives from it
	//you can put db connections into separate files & db helpers into another
	"go.mongodb.org/mongo-driver/v2/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
