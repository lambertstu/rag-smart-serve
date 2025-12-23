package database

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	globalClient *mongo.Client
	mu           sync.RWMutex

	clients   = make(map[string]*mongo.Client)
	clientsMu sync.RWMutex

	ErrNotFound = mongo.ErrNoDocuments
)

func Init(uri string, opts ...*options.ClientOptions) error {
	mu.Lock()
	defer mu.Unlock()

	// 如果已经初始化，先断开旧连接
	if globalClient != nil {
		_ = globalClient.Disconnect(context.Background())
	}

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

func ensureInitialized() (*mongo.Client, error) {
	mu.RLock()
	defer mu.RUnlock()
	if globalClient == nil {
		return nil, errors.New("global mongo client not initialized")
	}
	return globalClient, nil
}

func getClient(uri string) (*mongo.Client, error) {
	if uri == "" {
		return ensureInitialized()
	}

	clientsMu.RLock()
	if client, ok := clients[uri]; ok {
		clientsMu.RUnlock()
		return client, nil
	}
	clientsMu.RUnlock()

	clientsMu.Lock()
	defer clientsMu.Unlock()

	if client, ok := clients[uri]; ok {
		return client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo at %s: %w", uri, err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, fmt.Errorf("failed to ping mongo: %w", err)
	}

	clients[uri] = client
	return client, nil
}

func SetClient(client *mongo.Client) {
	mu.Lock()
	defer mu.Unlock()
	globalClient = client
}

func GetClient() *mongo.Client {
	mu.RLock()
	defer mu.RUnlock()
	return globalClient
}

func Close(ctx context.Context) error {
	var errs []error
	mu.Lock()
	if globalClient != nil {
		if err := globalClient.Disconnect(ctx); err != nil {
			errs = append(errs, err)
		}
		globalClient = nil
	}
	mu.Unlock()

	clientsMu.Lock()
	for uri, client := range clients {
		if err := client.Disconnect(ctx); err != nil {
			errs = append(errs, err)
		}
		delete(clients, uri)
	}
	clientsMu.Unlock()

	if len(errs) > 0 {
		return fmt.Errorf("failed to close some mongo clients: %v", errs)
	}
	return nil
}

type MongoManager[T any] struct {
	dbName         string
	collectionName string
	uri            string
}

func NewMongoManager[T any](uri, dbName, collectionName string) *MongoManager[T] {
	return &MongoManager[T]{
		uri:            uri,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (m *MongoManager[T]) getCollection() (*mongo.Collection, error) {
	client, err := getClient(m.uri)
	if err != nil {
		return nil, err
	}
	return client.Database(m.dbName).Collection(m.collectionName), nil
}

func (m *MongoManager[T]) InsertOne(ctx context.Context, doc *T) (*mongo.InsertOneResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.InsertOne(ctx, doc)
}

func (m *MongoManager[T]) InsertMany(ctx context.Context, docs []*T) (*mongo.InsertManyResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	// 将 []*T 转换为 []interface{}
	interfaceDocs := make([]interface{}, len(docs))
	for i, v := range docs {
		interfaceDocs[i] = v
	}
	return coll.InsertMany(ctx, interfaceDocs)
}

func (m *MongoManager[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	var result T
	err = coll.FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &result, nil
}

func (m *MongoManager[T]) FindMany(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*T, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	cursor, err := coll.Find(ctx, filter, opts...)
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

func (m *MongoManager[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.UpdateOne(ctx, filter, update, opts...)
}

func (m *MongoManager[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.UpdateMany(ctx, filter, update, opts...)
}

func (m *MongoManager[T]) ReplaceOne(ctx context.Context, filter interface{}, replacement *T, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.ReplaceOne(ctx, filter, replacement, opts...)
}

func (m *MongoManager[T]) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.DeleteOne(ctx, filter, opts...)
}

func (m *MongoManager[T]) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.DeleteMany(ctx, filter, opts...)
}

func (m *MongoManager[T]) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	coll, err := m.getCollection()
	if err != nil {
		return 0, err
	}
	return coll.CountDocuments(ctx, filter, opts...)
}

func (m *MongoManager[T]) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) ([]*T, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	cursor, err := coll.Aggregate(ctx, pipeline, opts...)
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

func (m *MongoManager[T]) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.BulkWrite(ctx, models, opts...)
}

// GetDatabase 获取数据库对象
func (m *MongoManager[T]) GetDatabase() (*mongo.Database, error) {
	client, err := getClient(m.uri)
	if err != nil {
		return nil, err
	}
	return client.Database(m.dbName), nil
}

// GetCollection 获取原生集合对象 (如果未初始化，会尝试初始化)
func (m *MongoManager[T]) GetCollection() (*mongo.Collection, error) {
	return m.getCollection()
}

func (m *MongoManager[T]) FindByID(ctx context.Context, id interface{}) (*T, error) {
	return m.FindOne(ctx, bson.M{"_id": id})
}

// CreateIndex 创建单个索引
// keys: 索引键，例如 bson.D{{"field1", 1}, {"field2", -1}}
// unique: 是否唯一索引
func (m *MongoManager[T]) CreateIndex(ctx context.Context, keys interface{}, unique bool) (string, error) {
	coll, err := m.getCollection()
	if err != nil {
		return "", err
	}
	model := mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}
	return coll.Indexes().CreateOne(ctx, model)
}

// CreateIndexes 批量创建索引
func (m *MongoManager[T]) CreateIndexes(ctx context.Context, models []mongo.IndexModel) ([]string, error) {
	coll, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	return coll.Indexes().CreateMany(ctx, models)
}
