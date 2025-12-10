package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// globalClient 保存默认的 MongoDB 客户端实例
	globalClient *mongo.Client

	ErrClientNotInit = errors.New("mongo client not initialized")
	ErrNotFound      = mongo.ErrNoDocuments
)

// Init 初始化全局 MongoDB 客户端
func Init(uri string, opts ...*options.ClientOptions) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	allOpts := append([]*options.ClientOptions{clientOpts}, opts...)

	client, err := mongo.Connect(ctx, allOpts...)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("failed to ping mongo: %w", err)
	}

	globalClient = client
	return nil
}

func SetClient(client *mongo.Client) {
	globalClient = client
}

func GetClient() *mongo.Client {
	return globalClient
}

func Close(ctx context.Context) error {
	if globalClient != nil {
		return globalClient.Disconnect(ctx)
	}
	return nil
}

// MongoManager 是一个用于 MongoDB 操作的通用管理器
type MongoManager[T any] struct {
	collection *mongo.Collection
}

// NewMongoManager 创建类型 T 的 MongoManager 新实例
func NewMongoManager[T any](dbName, collectionName string) *MongoManager[T] {
	if globalClient == nil {
		panic(ErrClientNotInit)
	}
	return &MongoManager[T]{
		collection: globalClient.Database(dbName).Collection(collectionName),
	}
}

// NewMongoManagerWithClient 使用特定客户端创建一个新实例
func NewMongoManagerWithClient[T any](client *mongo.Client, dbName, collectionName string) *MongoManager[T] {
	return &MongoManager[T]{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

// InsertOne 插入单个文档
func (m *MongoManager[T]) InsertOne(ctx context.Context, doc *T) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, doc)
}

// InsertMany 插入多个文档
func (m *MongoManager[T]) InsertMany(ctx context.Context, docs []*T) (*mongo.InsertManyResult, error) {
	// 将 []*T 转换为 []interface{}
	interfaceDocs := make([]interface{}, len(docs))
	for i, v := range docs {
		interfaceDocs[i] = v
	}
	return m.collection.InsertMany(ctx, interfaceDocs)
}

// FindOne 查找匹配过滤条件的单个文档
func (m *MongoManager[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	var result T
	err := m.collection.FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &result, nil
}

// FindMany 查找匹配过滤条件的多个文档
func (m *MongoManager[T]) FindMany(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*T, error) {
	cursor, err := m.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateOne 更新单个文档
func (m *MongoManager[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.collection.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany 更新多个文档
func (m *MongoManager[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.collection.UpdateMany(ctx, filter, update, opts...)
}

// ReplaceOne 替换单个文档
func (m *MongoManager[T]) ReplaceOne(ctx context.Context, filter interface{}, replacement *T, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return m.collection.ReplaceOne(ctx, filter, replacement, opts...)
}

// DeleteOne 删除单个文档
func (m *MongoManager[T]) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.collection.DeleteOne(ctx, filter, opts...)
}

// DeleteMany 删除多个文档
func (m *MongoManager[T]) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.collection.DeleteMany(ctx, filter, opts...)
}

// Count 返回匹配过滤条件的文档数量
func (m *MongoManager[T]) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return m.collection.CountDocuments(ctx, filter, opts...)
}

// Aggregate 执行聚合管道
func (m *MongoManager[T]) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) ([]*T, error) {
	cursor, err := m.collection.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*T
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// BulkWrite 执行通用的批量操作
func (m *MongoManager[T]) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return m.collection.BulkWrite(ctx, models, opts...)
}

// GetCollection 返回底层的 mongo.Collection 以供高级用法使用
func (m *MongoManager[T]) GetCollection() *mongo.Collection {
	return m.collection
}

// FindByID 是通过 _id 查找的辅助函数。它假设 _id 是 ObjectID 类型或 id 字段中提供的字符串。
func (m *MongoManager[T]) FindByID(ctx context.Context, id interface{}) (*T, error) {
	return m.FindOne(ctx, bson.M{"_id": id})
}
