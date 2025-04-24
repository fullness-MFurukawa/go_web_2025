package errortype

// 内部エラー型
type InternalError struct {
	message string
}

// errorインターフェイスの実装
func (e *InternalError) Error() string {
	return e.message
}

// コンストラクタ
func NewInternalError(message string) *InternalError {
	return &InternalError{message: message}
}
