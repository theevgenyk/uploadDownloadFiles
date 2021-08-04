package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type FilesInDB struct {
	ID         primitive.ObjectID `bson:"_id"`
	Length     int64              `bson:"length"`
	ChunkSize  int32              `bson:"chunkSize"`
	UploadDate time.Time          `bson:"uploadDate"`
	FileName   string             `bson:"filename"`
}
