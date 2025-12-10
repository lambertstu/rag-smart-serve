package etcd

import (
	"context"
	"testing"
	"time"
)

func TestEtcdClient(t *testing.T) {
	endpoints := []string{"localhost:2379"}
	client, err := NewClient(endpoints, 5*time.Second)
	if err != nil {
		t.Logf("Skipping test: failed to connect to etcd at %v: %v", endpoints, err)
		return
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 尝试 Put 一个测试 key 来验证连接是否真的可用
	testKey := "mongo"
	testValue := "mongodb://localhost:27030"
	////
	//err = client.PutValue(ctx, testKey, testValue)
	//if err != nil {
	//	t.Logf("Skipping test: failed to put value (etcd might not be running): %v", err)
	//	return
	//}

	val, err := client.GetValue(ctx, testKey)
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
		_ = client.PutValue(context.Background(), testKey, updateVal)
	}()

	// 监听变化
	watchChan := client.Watch(testKey)
	timeout := time.After(3 * time.Second)

	select {
	case watchResp := <-watchChan:
		found := false
		for _, event := range watchResp.Events {
			// event.Type == mvccpb.PUT
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
