package usecase

import "service-exercise/domain/model/categories"

// ユースケース:[商品カテゴリ一覧を閲覧する] インターフェイス
type CategoryList interface {
	Execute() ([]*categories.Category, error)
}
