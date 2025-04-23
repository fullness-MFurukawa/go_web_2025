package products

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"
)

// 商品単価を表す値オブジェクト
type ProductPrice struct {
	value int
}

// ゲッター
func (p ProductPrice) Value() int {
	return p.value
}

// Equatableインターフェイスの実装
func (p *ProductPrice) Equals(other model.Equatable) bool {

	otherProductPrice, ok := other.(*ProductPrice)
	if !ok {
		return false
	}

	return p.value == otherProductPrice.value
}

// コンストラクタ
func NewProductPrice(value int) (*ProductPrice, error) {
	// 商品単価の範囲確認
	if value < 50 || value > 9999 {
		return nil, errortype.NewDomainError("商品単価は、50以上10000未満でなければなりません。")
	}
	return &ProductPrice{value: value}, nil
}
