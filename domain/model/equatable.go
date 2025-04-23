package model

// 等価性を検証インターフェイス
type Equatable interface {

	// other Equatable: Equatableインタフェースを実装した構造体
	// 戻り値: true:等しい false:等しくない
	Equals(other Equatable) bool
}
