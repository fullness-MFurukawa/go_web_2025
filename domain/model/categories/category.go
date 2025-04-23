package categories

import (
	"fmt"
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"
)

// 商品カテゴリを表すエンティティ
type Category struct {
	id   *CategoryId   // カテゴリId
	name *CategoryName // カテゴリ名
}

// Categoryのidを返す
func (c *Category) Id() *CategoryId {
	return c.id
}

// Categoryのnameを返す
func (c *Category) Name() *CategoryName {
	return c.name
}

// Categoryのnameを変更する
func (c *Category) ChangeName(newName *CategoryName) error {
	if newName == nil {
		return errortype.NewDomainError("カテゴリ名は、必須です。")
	}
	c.name = newName
	return nil
}

// Equatableインターフェイスの実装
// 引数のインスタンスと等価かどうかを検証する
func (c *Category) Equals(other model.Equatable) bool {
	otherCategory, ok := other.(*Category)
	if !ok {
		return false
	}
	// 識別子の比較結果を返す
	return c.id.Equals(otherCategory.id)
}

// fmt.Stringerインターフェイスを	実装する
func (c *Category) String() string {
	var idStr, nameStr string
	if c.id != nil {
		idStr = c.id.Value() // idがnilでなければValueを呼び出す
	} else {
		idStr = "nil" // idがnilの場合は"nil"とする
	}
	if c.name != nil {
		nameStr = c.name.Value() // nameがnilでなければValueを呼び出す
	} else {
		nameStr = "nil" // nameがnilの場合は"nil"とする
	}
	return fmt.Sprintf("Category: id=%s,name=%s", idStr, nameStr)
}

// コンストラクタ Categoryエンティティを生成する
func NewCategory(id *CategoryId, name *CategoryName) (*Category, error) {
	if id == nil {
		return nil, errortype.NewDomainError("カテゴリIdは、必須です。")
	}
	return &Category{id: id, name: name}, nil
}
