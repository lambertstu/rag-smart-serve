package database

import (
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MetaCollection = "fs.files"
const ChunksCollection = "fs.chunks"

type GridFs struct {
	database *mongo.Database
}

func NewGridFs(database *mongo.Database) *GridFs {
	return &GridFs{database: database}
}

func (receiver *GridFs) Upload(filename string, reader io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error) {
	bucket, err := gridfs.NewBucket(receiver.database)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return bucket.UploadFromStream(filename, reader, opts...)
}

func (receiver *GridFs) Download(filename string) ([]byte, error) {
	bucket, err := gridfs.NewBucket(receiver.database)
	if err != nil {
		return nil, err
	}
	stream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	return io.ReadAll(stream)
}

func (receiver *GridFs) HandleByDownloadStream(filename string, handler func(downloadStream *gridfs.DownloadStream) error) error {
	bucket, err := gridfs.NewBucket(receiver.database)
	if err != nil {
		return err
	}
	stream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		return err
	}
	defer stream.Close()
	return handler(stream)
}
