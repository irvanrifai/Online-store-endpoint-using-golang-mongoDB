package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dimension struct {
	Length int32 `json:"length,omitempty" bson:"length,omitempty"`
	Width  int32 `json:"width,omitempty" bson:"width,omitempty"`
	Height int32 `json:"height,omitempty" bson:"height,omitempty"`
}

type Item struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Size      int64              `json:"size,omitempty" bson:"size,omitempty"`
	Dimension Dimension          `json:"dimension,omitempty" bson:"dimension,omitempty"`
	Price     string             `json:"price,omitempty" bson:"price,omitempty"`
	Quantity  int64              `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Desc      string             `json:"desc,omitempty" bson:"desc,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
