package connector

// データベース接続機能を実装するためのインターフェイス
type DatabaseConnector[T any] interface {
	// データベース接続メソッド
	Connect() (T, error)
}
