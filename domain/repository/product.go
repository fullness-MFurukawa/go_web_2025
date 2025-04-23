package repository

import "service-exercise/domain/model/products"

// 商品のCRUD操作リポジトリインターフェイス
type ProductRepository[T any] interface {
	// 指定された商品の存在確認
	Exists(db T, name *products.ProductName) (bool, error)
	// 新しい商品を永続化する
	Create(db T, product *products.Product) error
	// 商品をキーワード検索する
	FindByNameLike(db T, keyword string) ([]*products.Product, error)
}
