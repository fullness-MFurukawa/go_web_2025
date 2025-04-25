package adapter

import (
	"service-exercise/domain/adapter"
	"service-exercise/domain/model/categories"
	"service-exercise/presentation/dto"
)

// Categoryエンティティと他のモデルの相互変換インターフェイスの実装
type categoryAdapaterGin struct{}

// エンティティを他のモデルに変換する
func (c *categoryAdapaterGin) Convert(source *categories.Category) (*dto.CategoryDTO, error) {
	return dto.NewCatgeoryDTO(source.Id().Value(), source.Name().Value()), nil
}

// 他のモデルからエンティティを復元する
func (c *categoryAdapaterGin) Restore(source *dto.CategoryDTO) (*categories.Category, error) {
	categoryId, err := categories.NewCategoryId(source.Id)
	if err != nil {
		return nil, err
	}
	categoryName, err := categories.NewCategoryName(source.Name)
	if err != nil {
		return nil, err
	}
	category, err := categories.NewCategory(categoryId, categoryName)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// コンストラクタ
func NewcategoryAdapaterGin() adapter.CategoryAdapter[*dto.CategoryDTO] {
	return &categoryAdapaterGin{}
}
