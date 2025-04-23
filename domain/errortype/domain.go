package errortype

// ドメインエラー型
type DomainError struct {
	message string
}

// errorインターフェイスの実装
func (e *DomainError) Error() string {
	return e.message
}

// コンストラクタ
func NewDomainError(message string) *DomainError {
	return &DomainError{message: message}
}
