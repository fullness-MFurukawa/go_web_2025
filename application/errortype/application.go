package errortype

// アプリケーションエラー型
type ApplicationError struct {
	message string
}

// errorインターフェイスの実装
func (e *ApplicationError) Error() string {
	return e.message
}

// コンストラクタ
func NewApplicationError(message string) *ApplicationError {
	return &ApplicationError{message: message}
}
