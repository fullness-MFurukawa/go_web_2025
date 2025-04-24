package repimpl

import (
	"fmt"
	"service-exercise/domain/adapter"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/repository"
	"service-exercise/infrastructure/errortype"
	"service-exercise/infrastructure/gorm/dbmodel"

	"gorm.io/gorm"
)

// 商品カテゴリをCRUD操作するリポジトリインターフェイスの実装
type categtoryRepositoryImpl struct {
	// ドメイン層のCategoryAdapterインターフェース型
	adapter adapter.CategoryAdapter[*dbmodel.CategoryModel]
}

// すべての商品カテゴリを取得する
func (c *categtoryRepositoryImpl) FindAll(db *gorm.DB) ([]*categories.Category, error) {
	var models []dbmodel.CategoryModel
	// すべての商品カテゴリを取得してmodelsに格納する
	result := db.Find(&models)
	if result.Error != nil {
		return nil, errortype.NewInternalError(result.Error.Error())
	}
	// 取得したCategtoryModelをCategoryエンティティに変換する
	var categories []*categories.Category
	for _, model := range models {
		// CategoryModelからCategoryを復元する
		category, err := c.adapter.Restore(&model)
		if err != nil {
			return nil, errortype.NewInternalError(err.Error())
		}
		// categoriesスライスに復元したCategoryを追加する
		categories = append(categories, category)
	}
	return categories, nil
}

// fxによって実行される
// コンストラクタ
// コンストラクタインジェクションで利用したい、categoryAdapterImplを
// fxに入れてもらう
func NewcategtoryRepositoryImpl(
	// CategoryAdapterインターフェースを実装した構造体が渡せる
	adapter adapter.CategoryAdapter[*dbmodel.CategoryModel]) repository.CategtoryRepository[*gorm.DB] {
	fmt.Println("fxに生成された:categtoryRepositoryImpl")
	return &categtoryRepositoryImpl{adapter: adapter}
}
