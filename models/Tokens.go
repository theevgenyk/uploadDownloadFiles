package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	Value string `json:"token"`
}

type TokenFromDB struct {
	IdFile primitive.ObjectID `bson:"idFile"`
}
