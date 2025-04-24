package impls

import (
	"fmt"
	"service-exercise/application/errortype"
	"service-exercise/application/usecase"
	"service-exercise/domain/model/products"
	"service-exercise/domain/repository"

	"gorm.io/gorm"
)

// ユースケース:[商品を登録する] インターフェイスの実装
type productRegisterImpl struct {
	db *gorm.DB
	// 商品のCRUD操作リポジトリインターフェイス
	repository repository.ProductRepository[*gorm.DB]
}

// 商品を登録する
func (p *productRegisterImpl) Execute(newProduct *products.Product) (err error) {
	// トランザクションを開始する
	tx := p.db.Begin()
	if tx.Error != nil {
		err = tx.Error
		return
	}
	// 遅延関数でトランザクションを終了する
	defer func() {
		if p := recover(); p != nil { // パニックが発生?
			tx.Rollback()
			panic(p) // パニックを再発生させる
		} else if err != nil {
			tx.Rollback() // エラーがある場合はロールバック
		} else {
			err = tx.Commit().Error // エラーがない場合はコミット
		}
	}()

	// 登録しようとしいる商品が登録済みか確認する
	var exists bool
	if exists, err = p.repository.Exists(tx, newProduct.Name()); err != nil {
		return
	}
	if exists { // 既に登録されている
		err = errortype.NewApplicationError(
			fmt.Sprintf("商品:%sは、既に登録済です。", newProduct.Name().Value()))
		return
	}
	// 商品を登録する
	if err = p.repository.Create(tx, newProduct); err != nil {
		return
	}
	return nil
}

// コンストラクタ
func NewproductRegisterImpl(db *gorm.DB,
	repository repository.ProductRepository[*gorm.DB]) usecase.ProductRegister {
	return &productRegisterImpl{
		db:         db,
		repository: repository,
	}
}
