package impls

import (
	"service-exercise/application/usecase"
	"service-exercise/domain/model/products"
	"service-exercise/domain/repository"

	"gorm.io/gorm"
)

// ユースケース:[商品を検索する] インターフェイスの実装
type productKeywordImpl struct {
	db *gorm.DB
	// 商品のCRUD操作リポジトリインターフェイス
	repository repository.ProductRepository[*gorm.DB]
}

// 商品をキーワード検索する
func (e *productKeywordImpl) Execute(keyword string) ([]*products.Product, error) {
	return e.repository.FindByNameLike(e.db, keyword)
}

// コンストラクタ
func NewproductKeywordImpl(db *gorm.DB,
	repository repository.ProductRepository[*gorm.DB]) usecase.ProductKeyword {
	return &productKeywordImpl{
		db:         db,
		repository: repository,
	}
}
