package llm

import (
	"net/http"
)

func (c *ApiClient) Generate(req OllamaGenerateReq) (*OllamaGenerateResp, error) {
	return doRequest[OllamaGenerateResp](c, http.MethodPost, "/api/generate", req)
}

func (c *ApiClient) CreateEmbedding(req OllamaEmbeddingReq) (*OllamaEmbeddingResp, error) {
	return doRequest[OllamaEmbeddingResp](c, http.MethodPost, "/api/embeddings", req)
}
