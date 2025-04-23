package repository

import "service-exercise/domain/model/categories"

// 商品カテゴリをCRUD操作するリポジトリインターフェイス
type CategtoryRepository[T any] interface {
	// すべての商品カテゴリを取得する
	FindAll(db T) ([]*categories.Category, error)
}
