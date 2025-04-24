package repimpl

import (
	"fmt"
	"service-exercise/domain/adapter"
	"service-exercise/domain/model/products"
	"service-exercise/domain/repository"
	"service-exercise/infrastructure/errortype"
	"service-exercise/infrastructure/gorm/dbmodel"

	"gorm.io/gorm"
)

// 商品のCRUD操作リポジトリインターフェイスの実装
type productRepositoryImpl struct {
	adapter adapter.ProductAdapter[*dbmodel.ProductModel]
}

// 指定された商品の存在確認
func (p *productRepositoryImpl) Exists(db *gorm.DB, name *products.ProductName) (bool, error) {
	var product dbmodel.ProductModel
	result := db.First(&product, "name = ?", name.Value())
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	} else if result.Error != nil {
		return false, errortype.NewInternalError(result.Error.Error())
	} else {
		return true, nil
	}
}

// 新しい商品を永続化する
func (p *productRepositoryImpl) Create(db *gorm.DB, product *products.Product) error {
	// ProductエンティティをProductModelに変換する
	productModel, err := p.adapter.Convert(product)
	if err != nil {
		return err
	}
	// 商品を永続化する
	result := db.Create(productModel)
	if result.Error != nil {
		return errortype.NewInternalError(result.Error.Error())
	}
	return nil
}

// 商品をキーワード検索する
func (p *productRepositoryImpl) FindByNameLike(db *gorm.DB, keyword string) ([]*products.Product, error) {
	var models []dbmodel.ProductModel
	// nameカラムで部分一致検索
	result := db.Preload("Category").Where("name LIKE ?", "%"+keyword+"%").Find(&models)
	if result.Error != nil {
		return nil, errortype.NewInternalError(result.Error.Error())
	}
	if len(models) == 0 {
		return nil, errortype.NewNotFoundError(
			fmt.Sprintf("キーワード:'%s'に該当する商品は見つかりませんでした。", keyword))
	}
	// 取得したProductModelをProductエンティティに変換する
	var products []*products.Product
	for _, model := range models {
		// productModelからProductを復元する
		product, err := p.adapter.Restore(&model)
		if err != nil {
			return nil, errortype.NewInternalError(err.Error())
		}
		// productsスライスに変換したProductを追加する
		products = append(products, product)
	}
	return products, nil
}

// コンストラクタ
func NewproductRepositoryImpl(
	adapter adapter.ProductAdapter[*dbmodel.ProductModel]) repository.ProductRepository[*gorm.DB] {
	fmt.Println("fxによって生成:productRepositoryImpl")
	return &productRepositoryImpl{adapter: adapter}
}
