package data_service

import (
	"context"
	"io"
	"rag-smart-serve/pkg/database"
	dto "rag-smart-serve/pkg/database/dto/data-service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ FileModel = (*customFileModel)(nil)

type FileModel interface {
	Upload(filename string, reader io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error)
	Download(filename string) ([]byte, error)
	HandleByDownloadStream(filename string, handler func(stream *gridfs.DownloadStream) error) error
	FileExist(ctx context.Context, filename string) (bool, error)
}

type customFileModel struct {
	*database.MongoManager[dto.File]
	gridFs *database.GridFs
}

func NewFileModel(uri, db, collection string) (FileModel, error) {
	manager := database.NewMongoManager[dto.File](uri, db, collection)
	mongoDb, err := manager.GetDatabase()
	if err != nil {
		return nil, err
	}
	return &customFileModel{
		MongoManager: manager,
		gridFs:       database.NewGridFs(mongoDb),
	}, nil
}

func (m *customFileModel) Upload(filename string, reader io.Reader, opts ...*options.UploadOptions) (primitive.ObjectID, error) {
	return m.gridFs.Upload(filename, reader, opts...)
}

func (m *customFileModel) Download(filename string) ([]byte, error) {
	return m.gridFs.Download(filename)
}

func (m *customFileModel) HandleByDownloadStream(filename string, handler func(stream *gridfs.DownloadStream) error) error {
	return m.gridFs.HandleByDownloadStream(filename, handler)
}

func (m *customFileModel) FileExist(ctx context.Context, filename string) (bool, error) {
	count, err := m.Count(ctx, bson.M{"filename": filename})
	return count > 0, err
}
