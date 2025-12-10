package etcd

import (
	"context"
	"fmt"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// clientWrapper 封装了 etcd client v3
type clientWrapper struct {
	cli *clientv3.Client
	mu  sync.Mutex
}

var defaultClient = &clientWrapper{}

func Init(endpoints []string, timeout time.Duration) error {
	defaultClient.mu.Lock()
	defer defaultClient.mu.Unlock()

	if defaultClient.cli != nil {
		// 已经初始化过，关闭旧连接
		_ = defaultClient.cli.Close()
	}

	if timeout == 0 {
		timeout = 5 * time.Second
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	})
	if err != nil {
		return fmt.Errorf("failed to create etcd client: %w", err)
	}

	defaultClient.cli = cli
	return nil
}

func ensureInitialized() error {
	defaultClient.mu.Lock()
	defer defaultClient.mu.Unlock()

	if defaultClient.cli != nil {
		return nil
	}

	// 默认配置
	endpoints := []string{"localhost:2379"}
	timeout := 5 * time.Second

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	})
	if err != nil {
		return fmt.Errorf("failed to create default etcd client: %w", err)
	}

	defaultClient.cli = cli
	return nil
}

func GetValue(ctx context.Context, key string) (string, error) {
	if err := ensureInitialized(); err != nil {
		return "", err
	}

	resp, err := defaultClient.cli.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get value for key %s: %w", key, err)
	}

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("key %s not found", key)
	}

	return string(resp.Kvs[0].Value), nil
}

func PutValue(ctx context.Context, key, value string) error {
	if err := ensureInitialized(); err != nil {
		return err
	}

	_, err := defaultClient.cli.Put(ctx, key, value)
	if err != nil {
		return fmt.Errorf("failed to put value for key %s: %w", key, err)
	}
	return nil
}

func Watch(key string) clientv3.WatchChan {
	if err := ensureInitialized(); err != nil {
		// Watch 接口无法返回错误，如果初始化失败，只能 panic 或记录日志
		panic(fmt.Sprintf("failed to initialize default etcd client for Watch: %v", err))
	}
	return defaultClient.cli.Watch(context.Background(), key)
}

func Close() error {
	defaultClient.mu.Lock()
	defer defaultClient.mu.Unlock()

	if defaultClient.cli != nil {
		err := defaultClient.cli.Close()
		defaultClient.cli = nil
		return err
	}
	return nil
}
