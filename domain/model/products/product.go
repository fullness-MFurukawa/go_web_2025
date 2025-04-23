package products

import (
	"fmt"
	"service-exercise/domain/errortype"
	"service-exercise/domain/model"
	"service-exercise/domain/model/categories"
)

// 商品を表すエンティティ
type Product struct {
	id       *ProductId           // 商品Id
	name     *ProductName         // 商品名
	price    *ProductPrice        // 商品単価
	category *categories.Category // 商品カテゴリ
}

// Productのidを返す
func (p *Product) Id() *ProductId {
	return p.id
}

// Productのnameを返す
func (p *Product) Name() *ProductName {
	return p.name
}

// Productのpriceを返す
func (p *Product) Price() *ProductPrice {
	return p.price
}

// Productのcategoryを返す
func (p *Product) Category() *categories.Category {
	return p.category
}

// Productのnameを変更する
func (p *Product) ChangeName(newName *ProductName) error {
	if newName == nil {
		return errortype.NewDomainError("商品名は、必須です。")
	}
	p.name = newName
	return nil
}

// Productのpriceを変更する
func (p *Product) ChangePrice(newPrice *ProductPrice) error {
	if newPrice == nil {
		return errortype.NewDomainError("商品単価は、必須です。")
	}
	p.price = newPrice
	return nil
}

// Productのcategoryを変更する
func (p *Product) ChangeCategory(newCategory *categories.Category) error {
	if newCategory == nil {
		return errortype.NewDomainError("商品カテゴリは、必須です。")
	}
	p.category = newCategory
	return nil
}

// Equatableインターフェイスの実装
// 引数のインスタンスと等価かどうかを検証する
func (p *Product) Equals(other model.Equatable) bool {
	otherProduct, ok := other.(*Product)
	if !ok {
		return false
	}
	// 識別子の比較結果を返す
	return p.id.Equals(otherProduct.id)
}

// fmt.Stringerインターフェイスを	実装する
func (p *Product) String() string {
	var idStr, nameStr string
	if p.id != nil {
		idStr = p.id.Value() // idがnilでなければValueを呼び出す
	} else {
		idStr = "nil" // idがnilの場合は"nil"とする
	}
	if p.name != nil {
		nameStr = p.name.Value() // nameがnilでなければValueを呼び出す
	} else {
		nameStr = "nil" // nameがnilの場合は"nil"とする
	}
	return fmt.Sprintf("Product: id=%s,name=%s,price=%d,%s",
		idStr, nameStr, p.price.value, p.category.String())
}

// コンストラクタ Productエンティティを生成する
func NewProduct(id *ProductId, name *ProductName, price *ProductPrice,
	category *categories.Category) (*Product, error) {
	if id == nil {
		return nil, errortype.NewDomainError("商品Idは、必須です。")
	}
	return &Product{id: id, name: name, price: price, category: category}, nil
}
