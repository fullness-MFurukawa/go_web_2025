package repimpl

import (
	"fmt"
	"service-exercise/domain/adapter"
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"
	"service-exercise/infrastructure/gorm/dbmodel"
)

// 商品カテゴリエンティティと他のモデルの相互変換インターフェイスの実装
type categoryAdapterImpl struct{}

// エンティティを他のモデルに変換する
func (c *categoryAdapterImpl) Convert(source *categories.Category) (*dbmodel.CategoryModel, error) {
	if source == nil {
		return nil,
			errortype.NewDomainError("引数がnilのため、CategoryModelへの変換ができません。")
	}
	return &dbmodel.CategoryModel{
		ID:    0,
		ObjID: source.Id().Value(),
		Name:  source.Name().Value(),
	}, nil
}

// 他のモデルからエンティティを復元する
func (c *categoryAdapterImpl) Restore(source *dbmodel.CategoryModel) (*categories.Category, error) {
	if source == nil {
		return nil,
			errortype.NewDomainError("引数がnilのため、Categoryを復元できません。")
	}
	id, err := categories.NewCategoryId(source.ObjID)
	if err != nil {
		return nil, err
	}
	name, err := categories.NewCategoryName(source.Name)
	if err != nil {
		return nil, err
	}
	return categories.NewCategory(id, name)
}

// コンストラクタ
// fxによって実行される
func NewcategoryAdapterImpl() adapter.CategoryAdapter[*dbmodel.CategoryModel] {
	fmt.Println("fxに生成された：categoryAdapterImpl")
	return &categoryAdapterImpl{}
}
