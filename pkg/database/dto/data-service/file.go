package data_service

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Length     int                `bson:"length"`
	ChunkSize  int                `bson:"chunkSize"`
	UploadDate time.Time          `bson:"uploadDate"`
	Filename   string             `bson:"filename"`
	Metadata   map[string]any     `bson:"metadata"`
	UpdateAt   time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt   time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
