package dto

// 商品カテゴリDTO
type CategoryDTO struct {
	Id   string `json:"categoryId"`
	Name string `json:"categoryName"`
}

// コンストラクタ
func NewCatgeoryDTO(id string, name string) *CategoryDTO {
	return &CategoryDTO{Id: id, Name: name}
}
