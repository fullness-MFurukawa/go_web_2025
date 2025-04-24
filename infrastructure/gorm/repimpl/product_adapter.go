package repimpl

import (
	"fmt"
	"service-exercise/domain/adapter"
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/products"
	"service-exercise/infrastructure/gorm/dbmodel"
)

// 商品エンティティと他のモデルを相互変換するインターフェイスの実装
type productAdapterImpl struct {
	categoryAdapter adapter.CategoryAdapter[*dbmodel.CategoryModel]
}

// エンティティを他のモデルに変換する
func (p *productAdapterImpl) Convert(source *products.Product) (*dbmodel.ProductModel, error) {
	if source == nil {
		return nil, errortype.NewDomainError("引数がnilのため、ProductModelに変換できません。")
	}
	return &dbmodel.ProductModel{
		ID:         0,
		ObjID:      source.Id().Value(),
		Name:       source.Name().Value(),
		Price:      source.Price().Value(),
		CategoryID: source.Category().Id().Value(),
	}, nil
}

// 他のモデルからエンティティを復元する
func (p *productAdapterImpl) Restore(source *dbmodel.ProductModel) (*products.Product, error) {
	if source == nil {
		return nil, errortype.NewDomainError("引数がnilのため、Productを復元できません。")
	}
	id, err := products.NewProductId(source.ObjID)
	if err != nil {
		return nil, err
	}
	name, err := products.NewProductName(source.Name)
	if err != nil {
		return nil, err
	}
	price, err := products.NewProductPrice(source.Price)
	if err != nil {
		return nil, err
	}
	category, err := p.categoryAdapter.Restore(&source.Category)
	if err != nil {
		return nil, err
	}
	return products.NewProduct(id, name, price, category)
}

// コンストラクタ
func NewproductAdapterImpl(
	categoryAdapter adapter.CategoryAdapter[*dbmodel.CategoryModel]) adapter.ProductAdapter[*dbmodel.ProductModel] {
	fmt.Println("fxが生成:productAdapterImpl")
	return &productAdapterImpl{
		categoryAdapter: categoryAdapter,
	}
}
