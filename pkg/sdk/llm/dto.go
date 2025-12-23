package llm

import "fmt"

// BaseResponse 标准通用响应包装 (供未来使用)
// 示例: {"code": 0, "msg": "ok", "data": {...}}
type BaseResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

func (r *BaseResponse[T]) IsSuccess() bool {
	return r.Code == 0
}

func (r *BaseResponse[T]) GetError() error {
	if r.IsSuccess() {
		return nil
	}
	return fmt.Errorf("api biz error: code=%d, msg=%s", r.Code, r.Msg)
}

// ==================== Ollama API DTOs ====================

// OllamaGenerateReq 调用 /api/generate 接口的请求体
type OllamaGenerateReq struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaGenerateResp 是调用 /api/generate 接口的响应体
type OllamaGenerateResp struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"` // 业务 JSON 字符串
	Done      bool   `json:"done"`
	Context   []int  `json:"context,omitempty"`

	// 性能指标字段
	TotalDuration      int64 `json:"total_duration,omitempty"`
	LoadDuration       int64 `json:"load_duration,omitempty"`
	PromptEvalCount    int   `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64 `json:"prompt_eval_duration,omitempty"`
	EvalCount          int   `json:"eval_count,omitempty"`
	EvalDuration       int64 `json:"eval_duration,omitempty"`
}

// OllamaEmbeddingReq 调用 /api/embeddings 接口的请求体
type OllamaEmbeddingReq struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// OllamaEmbeddingResp 调用 /api/embeddings 接口的响应体
type OllamaEmbeddingResp struct {
	Embedding []float64 `json:"embedding"`
}

// EmotionResult 业务层实体 (从 Response 字符串中解析)
type EmotionResult struct {
	Type       string  `json:"type"`       // positive, neutral, negative, angry
	Score      float64 `json:"score"`      // 0.0 - 1.0
	RiskLevel  int     `json:"risk_level"` // 0 - 3
	Suggestion string  `json:"suggestion"` // pass, comfort, transfer_human, reject
	Reason     string  `json:"reason"`
}

type EmotionOutput struct {
	Reason     string  `json:"reason"`
	RiskLevel  int     `json:"risk_level"`
	Score      float64 `json:"score"`
	Suggestion string  `json:"suggestion"`
	Type       string  `json:"type"`
}
