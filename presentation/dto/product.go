package dto

// 商品DTO
type ProductDTO struct {
	Id       string       `json:"productId"`
	Name     string       `json:"productName"`
	Price    string       `json:"productPrice"`
	Category *CategoryDTO `json:"category,omitempty"`
}

// コンストラクタ
func NewProductDTO(id string, name string, price string, dto *CategoryDTO) *ProductDTO {
	return &ProductDTO{Id: id, Name: name, Price: price, Category: dto}
}
