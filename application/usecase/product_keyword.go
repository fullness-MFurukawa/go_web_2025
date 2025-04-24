package usecase

import "service-exercise/domain/model/products"

// ユースケース:[商品を検索する] インターフェイス
type ProductKeyword interface {
	// keyword:画面に入力された商品キーワード
	Execute(keyword string) ([]*products.Product, error)
}
