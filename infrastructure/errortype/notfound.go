package errortype

// レコードが見つからないエラー型
type NotFoundError struct {
	message string
}

// errorインターフェイスの実装
func (e *NotFoundError) Error() string {
	return e.message
}

// コンストラクタ
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message: message}
}
