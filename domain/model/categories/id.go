package categories

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"

	"github.com/google/uuid"
)

// 商品カテゴリIdを表す値オブジェクト
type CategoryId struct {
	value string
}

// ゲッター
// 外部から商品カテゴリIdを読み取るためのメソッド
func (c *CategoryId) Value() string {
	return c.value
}

// Equatableインターフェイスの実装
func (c *CategoryId) Equals(other model.Equatable) bool {
	// 型アサーション
	// 引数otherはCategoryId型かを調べる
	otherCategoryId, ok := other.(*CategoryId)
	if !ok {
		return false // CategoryId型でなければfalseを返す
	}
	// 引数で渡されたCategoryIdのvalueフィールドと等価であるかを検証
	// 同じ値ならtrue 違う値ならfalseを返す
	return c.value == otherCategoryId.value
}

// 新しいUUIDを生成し、その値を持つCategoryIdを返す
// 新しい商品カテゴリIdを作りたい
func NewCategoryIdWithUUID() *CategoryId {
	categoryId := CategoryId{value: uuid.NewString()}
	return &categoryId
}

// コンストラクタ
// 既に存在しているUUIDから商品カテゴリIdを生成する
func NewCategoryId(value string) (*CategoryId, error) {
	if value == "" {
		return nil, errortype.NewDomainError("商品カテゴリIdは、空文字列であってはなりません。")
	}
	if len(value) != 36 {
		return nil, errortype.NewDomainError("商品カテゴリIdは、36文字でなければなりません。")
	}
	if _, err := uuid.Parse(value); err != nil {
		return nil, errortype.NewDomainError("商品カテゴリIdは、UUID形式でなければなりません。")
	}
	categoryId := CategoryId{value: value}
	return &categoryId, nil
}
