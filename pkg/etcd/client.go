package etcd

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// Client 封装了 etcd client v3
type Client struct {
	cli *clientv3.Client
}

// NewClient 初始化一个新的 etcd 客户端
// endpoints: etcd 地址列表，例如 []string{"localhost:2379"}
// timeout: 连接超时时间
func NewClient(endpoints []string, timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %w", err)
	}

	return &Client{cli: cli}, nil
}

// GetValue 获取指定 key 的值
func (c *Client) GetValue(ctx context.Context, key string) (string, error) {
	resp, err := c.cli.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("failed to get value for key %s: %w", key, err)
	}

	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf("key %s not found", key)
	}

	return string(resp.Kvs[0].Value), nil
}

// PutValue 设置 key-value
func (c *Client) PutValue(ctx context.Context, key, value string) error {
	_, err := c.cli.Put(ctx, key, value)
	if err != nil {
		return fmt.Errorf("failed to put value for key %s: %w", key, err)
	}
	return nil
}

// Watch 监听 key 的变化
func (c *Client) Watch(key string) clientv3.WatchChan {
	return c.cli.Watch(context.Background(), key)
}

// Close 关闭客户端连接
func (c *Client) Close() error {
	return c.cli.Close()
}
