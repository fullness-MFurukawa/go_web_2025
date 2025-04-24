package dbmodel

// 商品用GORMモデル
type ProductModel struct {
	ID uint `gorm:"primaryKey"`
	// 商品Id
	ObjID string `gorm:"uniqueIndex;type:varchar(36)"`
	// 商品名
	Name string `gorm:"type:varchar(30)"`
	// 単価
	Price int
	// カテゴリId(外部キー)
	CategoryID string `gorm:"type:varchar(36)"`
	// 商品カテゴリ(結合結果)
	Category CategoryModel `gorm:"foreignKey:CategoryID;references:ObjID"`
}

// GORMにこのモデルが使用するテーブル名を教える
func (ProductModel) TableName() string {
	return "product"
}
