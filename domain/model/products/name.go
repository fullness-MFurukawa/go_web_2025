package products

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"
	"unicode/utf8"
)

// 商品名を表す値オブジェクト
type ProductName struct {
	value string
}

// ゲッター
func (p *ProductName) Value() string {
	return p.value
}

// Equatableインターフェイスの実装
func (p *ProductName) Equals(other model.Equatable) bool {

	otherProductName, ok := other.(*ProductName)
	if !ok {
		return false
	}

	return p.value == otherProductName.value
}

// コンストラクタ
func NewProductName(value string) (*ProductName, error) {
	if value == "" {
		return nil, errortype.NewDomainError("商品名は、空文字列であってはなりません。")
	}
	if utf8.RuneCountInString(value) > 30 {
		return nil, errortype.NewDomainError("商品名は、30文字以内である必要があります。")
	}
	productName := ProductName{value: value}
	return &productName, nil
}
