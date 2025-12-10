package sdk

import (
	"rag-smart-serve/pkg/constant"
	"testing"
)

func TestApiClient_Generate(t *testing.T) {
	client := NewApiClient("http://localhost:11434")

	req := OllamaGenerateReq{
		Model:  constant.EmotionModel,
		Prompt: "请问这个商品款号是什么",
		Stream: false,
	}

	resp, err := client.Generate(req)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	t.Logf("Raw Response: %s", resp.Response)
}

func TestApiClient_Embedding(t *testing.T) {
	client := NewApiClient("http://localhost:11434")

	req := OllamaEmbeddingReq{
		Model:  constant.EmbeddingModel,
		Prompt: "示例demo",
	}
	resp, err := client.CreateEmbedding(req)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	t.Logf("Raw Response: %f", resp)
}
