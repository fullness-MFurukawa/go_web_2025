package usecase

import "service-exercise/domain/model/products"

// ユースケース:[商品を登録する] インターフェイス
type ProductRegister interface {
	// newProductは画面で入力された商品名や単価を格納しているエンティティ
	Execute(newProduct *products.Product) (err error)
}
