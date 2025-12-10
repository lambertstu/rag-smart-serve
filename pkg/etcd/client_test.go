package etcd

import (
	"context"
	"testing"
	"time"
)

func TestEtcdClient(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// 尝试 Put 一个测试 key
	testKey := "mongo"
	testValue := "mongodb://localhost:27030"

	val, err := GetValue(ctx, testKey)
	if err != nil {
		t.Fatalf("GetValue failed: %v", err)
	}
	if val != testValue {
		t.Errorf("Expected value %s, got %s", testValue, val)
	}

	// 3. 测试 Watch
	// 启动一个 goroutine 来修改值
	go func() {
		time.Sleep(1 * time.Second)
		updateVal := "mongodb://localhost:27030/?directConnection=true"
		_ = PutValue(context.Background(), testKey, updateVal)
	}()

	// 监听变化
	watchChan := Watch(testKey)
	timeout := time.After(3 * time.Second)

	select {
	case watchResp := <-watchChan:
		found := false
		for _, event := range watchResp.Events {
			if string(event.Kv.Value) == "mongodb://localhost:27030/?directConnection=true" {
				found = true
				break
			}
		}
		if !found {
			t.Error("Watch received event but value did not match expected update")
		}
	case <-timeout:
		t.Error("Watch timed out waiting for update")
	}
}
