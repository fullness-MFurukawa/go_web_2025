package categories

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"
	"unicode/utf8"
)

// 商品カテゴリを表す値オブジェクト
type CategoryName struct {
	value string
}

// ゲッター
func (c *CategoryName) Value() string {
	return c.value
}

// Equatableインターフェイスの実装
func (c *CategoryName) Equals(other model.Equatable) bool {

	otherCategoryName, ok := other.(*CategoryName)
	if !ok {
		return false
	}
	return c.value == otherCategoryName.value
}

// コンストラクタ
func NewCategoryName(value string) (*CategoryName, error) {
	if value == "" {
		return nil, errortype.NewDomainError("商品カテゴリ名は、空文字列であってはなりません。")
	}
	if utf8.RuneCountInString(value) > 20 {
		return nil, errortype.NewDomainError("商品カテゴリ名は、20文字以内である必要があります。")
	}
	categoryName := CategoryName{value: value}
	return &categoryName, nil
}
