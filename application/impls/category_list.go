package impls

import (
	"service-exercise/application/usecase"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/repository"

	"gorm.io/gorm"
)

// ユースケース:[商品カテゴリ一覧を閲覧する] インターフェイスの実装
type categoryListImpl struct {
	db *gorm.DB
	// 商品カテゴリをCRUD操作するリポジトリインターフェイス
	repository repository.CategtoryRepository[*gorm.DB]
}

// 　すべてのカテゴリを取得する
func (c *categoryListImpl) Execute() ([]*categories.Category, error) {
	return c.repository.FindAll(c.db)
}

// コンストラクタ
func NewcategoryListImpl(db *gorm.DB,
	repository repository.CategtoryRepository[*gorm.DB]) usecase.CategoryList {
	return &categoryListImpl{
		db:         db,
		repository: repository,
	}
}
