package repimpl

import (
	"fmt"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/repository"

	"gorm.io/gorm"
)

type categoryRepositoryImpl2 struct {
}

// すべての商品カテゴリを取得する
func (c *categoryRepositoryImpl2) FindAll(db *gorm.DB) ([]*categories.Category, error) {
	fmt.Println("**********************************************************")
	fmt.Println("************* CategoryRepositoryインターフェースの実装2 ****")
	fmt.Println("**********************************************************")
	return nil, nil
}

// コンストラクタ
func NewcategoryRepositoryImpl2() repository.CategtoryRepository[*gorm.DB] {
	return &categoryRepositoryImpl2{}
}
