package database

import (
	"context"
	"rag-smart-serve/pkg/constant"
	"testing"
	"time"

	"rag-smart-serve/pkg/etcd"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	CreatedAt time.Time          `bson:"created_at"`
	Age       int                `bson:"age"`
}

func TestMongoManager(t *testing.T) {
	endpoints := []string{"localhost:2379"}
	client, err := etcd.NewClient(endpoints, 5*time.Second)
	if err != nil {
		t.Logf("Skipping test: failed to connect to etcd at %v: %v", endpoints, err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	uri, err := client.GetValue(ctx, constant.MongoKey)
	err = Init(uri)
	if err != nil {
		t.Logf("Skipping test: failed to connect to mongo at %s: %v", uri, err)
		return
	}
	defer Close(context.Background())

	// 2. 创建 Manager
	dbName := "test_db"
	collName := "users"
	manager := NewMongoManager[TestUser](dbName, collName)

	_ = manager.GetCollection().Drop(ctx)

	// 3. 测试插入 (InsertOne)
	user1 := &TestUser{
		Username:  "alice",
		Email:     "alice@example.com",
		Age:       25,
		CreatedAt: time.Now(),
	}
	insertRes, err := manager.InsertOne(ctx, user1)
	if err != nil {
		t.Fatalf("InsertOne failed: %v", err)
	}
	if insertRes.InsertedID == nil {
		t.Fatal("InsertOne returned nil ID")
	}
	// 将 ID 回填方便后续查询
	user1.ID = insertRes.InsertedID.(primitive.ObjectID)

	// 4. 测试查询单个 (FindOne)
	foundUser, err := manager.FindByID(ctx, user1.ID)
	if err != nil {
		t.Fatalf("FindByID failed: %v", err)
	}
	if foundUser.Username != user1.Username {
		t.Errorf("Expected username %s, got %s", user1.Username, foundUser.Username)
	}

	// 5. 测试插入多个 (InsertMany)
	users := []*TestUser{
		{Username: "bob", Age: 30, CreatedAt: time.Now()},
		{Username: "charlie", Age: 35, CreatedAt: time.Now()},
	}
	_, err = manager.InsertMany(ctx, users)
	if err != nil {
		t.Fatalf("InsertMany failed: %v", err)
	}

	// 6. 测试查询列表 (FindMany)
	// 查找年龄 >= 30 的用户
	filter := bson.M{"age": bson.M{"$gte": 30}}
	results, err := manager.FindMany(ctx, filter)
	if err != nil {
		t.Fatalf("FindMany failed: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("Expected 2 users, got %d", len(results))
	}

	// 7. 测试更新 (UpdateOne)
	updateFilter := bson.M{"username": "alice"}
	update := bson.M{"$set": bson.M{"age": 26}}
	_, err = manager.UpdateOne(ctx, updateFilter, update)
	if err != nil {
		t.Fatalf("UpdateOne failed: %v", err)
	}

	updatedAlice, err := manager.FindOne(ctx, updateFilter)
	if err != nil {
		t.Fatalf("FindOne after update failed: %v", err)
	}
	if updatedAlice.Age != 26 {
		t.Errorf("Expected age 26, got %d", updatedAlice.Age)
	}

	// 8. 测试删除 (DeleteOne)
	_, err = manager.DeleteOne(ctx, bson.M{"username": "bob"})
	if err != nil {
		t.Fatalf("DeleteOne failed: %v", err)
	}

	count, err := manager.Count(ctx, bson.M{})
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	// alice (1) + charlie (1) = 2
	if count != 2 {
		t.Errorf("Expected 2 documents remaining, got %d", count)
	}
}
