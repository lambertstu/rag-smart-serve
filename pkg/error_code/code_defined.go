package error_code

// ============================================================================
// 通用错误码
// 结构: [A/B/C] + 00(通用模块) + xxxx
// ============================================================================

var (
	// A: 客户端错误
	ClientError           = NewBaseErrorCode("A000001", "客户端请求错误")
	ParamVerifyError      = NewBaseErrorCode("A001000", "参数校验失败")
	AuthCheckError        = NewBaseErrorCode("A001010", "未登录或Token无效")
	PermissionDeniedError = NewBaseErrorCode("A001011", "无权限执行此操作")
	ResourceNotFoundError = NewBaseErrorCode("A001020", "请求资源不存在")
	RequestLimitError     = NewBaseErrorCode("A001030", "请求过于频繁，请稍后重试")

	// B: 服务端通用错误
	ServiceError        = NewBaseErrorCode("B000001", "系统内部错误")
	ServiceTimeoutError = NewBaseErrorCode("B001000", "系统执行超时")
	ServiceBusyError    = NewBaseErrorCode("B001010", "系统繁忙")
	DatabaseError       = NewBaseErrorCode("B001020", "数据库操作异常")
	CacheError          = NewBaseErrorCode("B001021", "缓存操作异常")

	// C: 第三方调用错误
	RemoteServiceError = NewBaseErrorCode("C000001", "调用第三方服务失败")
	RemoteTimeoutError = NewBaseErrorCode("C001000", "第三方服务响应超时")
)

// ============================================================================
// 网关与认证服务 (Gateway & Auth) - B01xxx / B02xxx
// ============================================================================
var (
	// B01: API Gateway
	GatewayRouteError = NewBaseErrorCode("B010001", "网关路由转发失败")

	// B02: Auth Service
	UserLoginFailed   = NewBaseErrorCode("B020001", "用户登录失败")
	TokenGenFailed    = NewBaseErrorCode("B020002", "Token生成失败")
	UserNotFound      = NewBaseErrorCode("B020003", "用户不存在")
	PasswordIncorrect = NewBaseErrorCode("B020004", "密码错误")
)

// ============================================================================
// 核心业务服务 (Core Services)
// ============================================================================

// RAG Service (B03)
var (
	RagRetrieveError  = NewBaseErrorCode("B030001", "知识库检索失败")
	RagGenerateError  = NewBaseErrorCode("B030002", "回复生成失败")
	RagContextTooLong = NewBaseErrorCode("B030003", "上下文长度超出限制")
)

// Intent Service (B04)
var (
	IntentDetectError = NewBaseErrorCode("B040001", "意图识别失败")
	SlotExtractError  = NewBaseErrorCode("B040002", "槽位提取失败")
)

// Emotion Service (B05)
var (
	EmotionAnalyzeError     = NewBaseErrorCode("B050001", "情绪分析失败")
	EmotionOutputError      = NewBaseErrorCode("B050002", "情绪分析输出结果不为标准值")
	ContentModerationFailed = NewBaseErrorCode("B050003", "内容审核未通过") // 内容违规
	RiskDetectError         = NewBaseErrorCode("B050004", "风险检测服务异常")
)

// Routing Service (B06)
var (
	RoutingDecisionError = NewBaseErrorCode("B060001", "路由决策失败")
	TransferHumanFailed  = NewBaseErrorCode("B060002", "转接人工失败")
	NoAgentAvailable     = NewBaseErrorCode("B060003", "当前无人工客服在线")
)

// Prompt Service (B07)
var (
	PromptTemplateNotFound = NewBaseErrorCode("B070001", "Prompt模板不存在")
	PromptRenderError      = NewBaseErrorCode("B070002", "Prompt模板渲染失败")
)

// ============================================================================
// 数据服务 (Data Services)
// ============================================================================

// Knowledge Service (B08)
var (
	KBCreateError       = NewBaseErrorCode("B080001", "知识库创建失败")
	DocUploadError      = NewBaseErrorCode("B080002", "文档上传失败")
	DocParseError       = NewBaseErrorCode("B080003", "文档解析失败")
	VectorIndexError    = NewBaseErrorCode("B080004", "向量索引构建失败")
	EmbeddingGenError   = NewBaseErrorCode("B080005", "Embedding生成失败")
	VectorDatabaseError = NewBaseErrorCode("B080006", "向量数据库操作失败")
)

// Analytics Service (B09)
var (
	ReportGenError = NewBaseErrorCode("B090001", "报表生成失败")
	LogRecordError = NewBaseErrorCode("B090002", "日志记录失败")
)

// ============================================================================
// 外部依赖错误 (External/Remote) - Cxxxxx
// ============================================================================

var (
	LLMServiceError      = NewBaseErrorCode("C001001", "LLM模型服务异常")
	LLMQuotaExceeded     = NewBaseErrorCode("C001002", "LLM调用配额不足")
	VectorDBServiceError = NewBaseErrorCode("C002001", "向量数据库服务不可用")
	RedisServiceError    = NewBaseErrorCode("C003001", "Redis服务异常")
	OSSStorageError      = NewBaseErrorCode("C004001", "对象存储服务异常")
)
