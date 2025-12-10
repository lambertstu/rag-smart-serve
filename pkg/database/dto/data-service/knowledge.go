package data_service

import "time"

type KnowledgeDao struct {
	Name           string    `json:"name" bson:"name"`
	Description    string    `json:"description" bson:"description"`
	EmbeddingModel string    `json:"embedding_model" bson:"embeddingModel"`
	CreatedAt      time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updatedAt"`
}
