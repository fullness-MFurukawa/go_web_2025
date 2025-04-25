package adapter

import (
	"fmt"
	"service-exercise/domain/adapter"
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/products"
	"service-exercise/presentation/dto"

	"strconv"
)

// 商品エンティティと他のモデルを相互変換するインターフェイスの実装
type productAdapterGin struct {
	adapter adapter.CategoryAdapter[*dto.CategoryDTO]
}

// エンティティから他のモデルへ変換する
func (p *productAdapterGin) Convert(source *products.Product) (*dto.ProductDTO, error) {
	// 商品カテゴリDTOを生成する
	category, _ := p.adapter.Convert(source.Category())
	// 商品DTOを生成する
	return dto.NewProductDTO(
		source.Id().Value(),
		source.Name().Value(),
		fmt.Sprint(source.Price().Value()),
		category), nil
}

// 他のモデルからエンティティを復元する
func (p *productAdapterGin) Restore(source *dto.ProductDTO) (*products.Product, error) {
	category, err := p.adapter.Restore(source.Category)
	if err != nil {
		return nil, err
	}
	var productId *products.ProductId
	if source.Id == "" {
		productId = products.NewProductIdWithUUID()
	} else {
		productId, err = products.NewProductId(source.Id)
		if err != nil {
			return nil, err
		}
	}
	productName, err := products.NewProductName(source.Name)
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(source.Price)
	if err != nil {
		return nil, errortype.NewDomainError("単価は整数でなければなりません。")
	}
	productPrice, err := products.NewProductPrice(price)
	if err != nil {
		return nil, err
	}
	return products.NewProduct(productId, productName, productPrice, category)
}

// コンストラクタ
func NewproductAdapterGin(
	adapter adapter.CategoryAdapter[*dto.CategoryDTO]) adapter.ProductAdapter[*dto.ProductDTO] {
	return &productAdapterGin{
		adapter: adapter,
	}
}
