package emotionservicelogic

import (
	"context"
	"core-service/core"
	"core-service/internal/svc"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"rag-smart-serve/pkg/constant"
	customErr "rag-smart-serve/pkg/error_code"
	"rag-smart-serve/pkg/sdk/llm"
)

type AnalyzeEmotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalyzeEmotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalyzeEmotionLogic {
	return &AnalyzeEmotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分析用户情绪
func (receiver *AnalyzeEmotionLogic) AnalyzeEmotion(in *core.AnalyzeEmotionReq) (*core.AnalyzeEmotionResp, error) {
	resp, err := receiver.svcCtx.ApiClient.Generate(
		llm.OllamaGenerateReq{
			Model:  constant.EmotionModel,
			Prompt: in.UserInput,
			Stream: false,
		})
	if err != nil {
		return nil, customErr.EmotionAnalyzeError.WithError(err)
	}

	var output llm.EmotionOutput
	err = json.Unmarshal([]byte(resp.Response), &output)
	if err != nil {
		return nil, customErr.EmotionOutputError
	}

	return &core.AnalyzeEmotionResp{
		EmotionType: output.Type,
		Score:       float32(output.Score),
		RiskLevel:   int32(output.RiskLevel),
		Suggestion:  output.Suggestion,
		Reason:      output.Reason,
	}, nil
}
