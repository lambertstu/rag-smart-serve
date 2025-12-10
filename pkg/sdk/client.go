package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ApiClient 通用 HTTP 客户端
type ApiClient struct {
	EndPoint   string
	Timeout    time.Duration
	HttpClient *http.Client
}

// NewApiClient 创建默认客户端
func NewApiClient(endPoint string) *ApiClient {
	return &ApiClient{
		EndPoint: endPoint,
		Timeout:  30 * time.Second,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// SetEndPoint 设置服务地址
func (c *ApiClient) SetEndPoint(endPoint string) *ApiClient {
	c.EndPoint = endPoint
	return c
}

// SetTimeout 设置超时时间
func (c *ApiClient) SetTimeout(d time.Duration) *ApiClient {
	c.Timeout = d
	c.HttpClient.Timeout = d
	return c
}

// doRequest 处理通用请求
// T: 响应体的结构类型
func doRequest[T any](c *ApiClient, method, requestUri string, requestData interface{}) (*T, error) {
	var bodyReader io.Reader

	if requestData != nil {
		reqPayload, err := json.Marshal(requestData)
		if err != nil {
			return nil, fmt.Errorf("marshal request failed: %w", err)
		}
		bodyReader = bytes.NewReader(reqPayload)
	}

	fullURL := c.EndPoint + requestUri
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// 这里可以添加通用的 Header，比如 Auth Token
	// req.Header.Set("Authorization", "Bearer xxx")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: status=%d body=%s", resp.StatusCode, string(body))
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	return &result, nil
}

// doGetRequest 处理 GET 请求（自动转换 Query 参数）
func doGetRequest[T any](c *ApiClient, requestUri string, requestData interface{}) (*T, error) {
	queryParams := buildQueryParams(requestData)
	fullUrl := c.EndPoint + requestUri
	if queryParams != "" {
		if strings.Contains(fullUrl, "?") {
			fullUrl += "&" + queryParams
		} else {
			fullUrl += "?" + queryParams
		}
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api error: status=%d body=%s", resp.StatusCode, string(body))
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("unmarshal response failed: %w", err)
	}

	return &result, nil
}

// buildQueryParams 将结构体转换为 URL Query String (参考您的示例代码)
func buildQueryParams(data interface{}) string {
	if data == nil {
		return ""
	}

	params := url.Values{}
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return ""
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		typeField := typ.Field(i)

		tag := typeField.Tag.Get("form") // 兼容 gin/go-zero 风格 tag
		if tag == "" {
			tag = typeField.Tag.Get("json") // 如果没有 form tag，尝试 json tag
		}
		if tag == "" || tag == "-" {
			continue
		}

		tagParts := strings.Split(tag, ",")
		fieldName := tagParts[0]

		if !field.IsValid() || (field.Kind() == reflect.Ptr && field.IsNil()) {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				params.Add(fieldName, field.String())
			}
		case reflect.Int, reflect.Int64:
			if field.Int() != 0 {
				params.Add(fieldName, strconv.FormatInt(field.Int(), 10))
			}
		case reflect.Bool:
			params.Add(fieldName, strconv.FormatBool(field.Bool()))
		case reflect.Float64:
			if field.Float() != 0 {
				params.Add(fieldName, strconv.FormatFloat(field.Float(), 'f', -1, 64))
			}
		}
	}
	return params.Encode()
}
