package error_code

import "fmt"

// ErrorCode 定义错误码行为接口
type ErrorCode interface {
	Code() string
	Message() string
	Details() string
	Error() string
}

// BaseErrorCode 基础错误码实现
type BaseErrorCode struct {
	code    string
	message string
	details string // 具体的错误详情，例如：哪个参数校验失败
}

// NewBaseErrorCode 创建新的错误码实例
func NewBaseErrorCode(code, message string) *BaseErrorCode {
	return &BaseErrorCode{
		code:    code,
		message: message,
	}
}

func (e *BaseErrorCode) Code() string {
	return e.code
}

func (e *BaseErrorCode) Message() string {
	return e.message
}

func (e *BaseErrorCode) Details() string {
	return e.details
}

// Error 实现 error 接口，返回格式化的错误信息
func (e *BaseErrorCode) Error() string {
	if e.details != "" {
		return fmt.Sprintf("[%s] %s - %s", e.code, e.message, e.details)
	}
	return fmt.Sprintf("[%s] %s", e.code, e.message)
}

// WithDetails 附加具体的错误详情 (返回新实例，不修改原错误码)
func (e *BaseErrorCode) WithDetails(details string) *BaseErrorCode {
	return &BaseErrorCode{
		code:    e.code,
		message: e.message,
		details: details,
	}
}

// WithError 包装底层 error
func (e *BaseErrorCode) WithError(err error) *BaseErrorCode {
	if err == nil {
		return e
	}
	return e.WithDetails(err.Error())
}

// ToError 用于需要显式返回 error 接口的场景
func (e *BaseErrorCode) ToError() error {
	return e
}
