package products

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"

	"github.com/google/uuid"
)

// 商品Idを表す値オブジェクト
type ProductId struct {
	value string
}

// 　ゲッター
func (p *ProductId) Value() string {
	return p.value
}

// Equatableインターフェイスの実装
func (p *ProductId) Equals(other model.Equatable) bool {

	otherProductId, ok := other.(*ProductId)
	if !ok {
		return false
	}

	return p.value == otherProductId.value
}

// 新しいUUIDを生成し、その値を持つCategoryIdを返す
func NewProductIdWithUUID() *ProductId {
	categoryId := ProductId{value: uuid.NewString()}
	return &categoryId
}

// コンストラクタ
func NewProductId(value string) (*ProductId, error) {
	if value == "" {
		return nil, errortype.NewDomainError("商品Idは、空文字列であってはなりません。")
	}
	if len(value) != 36 {
		return nil, errortype.NewDomainError("商品Idは、36文字でなければなりません。")
	}
	if _, err := uuid.Parse(value); err != nil {
		return nil, errortype.NewDomainError("商品Idは、UUID形式でなければなりません。")
	}
	productId := ProductId{value: value}
	return &productId, nil
}
