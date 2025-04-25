package handler

import (
	"net/http"
	"service-exercise/application/usecase"
	"service-exercise/domain/adapter"
	"service-exercise/presentation/dto"

	"github.com/gin-gonic/gin"
)

// 商品カテゴリ一覧を返すリクエストハンドラ
type categoryListHandler struct {
	usecase usecase.CategoryList
	adapter adapter.CategoryAdapter[*dto.CategoryDTO]
}

// 商品カテゴリ一覧を取得する
// @Tags 商品カテゴリ
// @Summary 商品カテゴリ一覧を取得する
// @Description 商品カテゴリ一覧を取得する
// @Id list-category
// @Accept application/json
// @Produce json
// @Success 200 {object} []dto.CategoryDTO
// @Failure 500 {object} errortype.InternalError
// @Router /category/list [GET]
func (c *categoryListHandler) Execute(context *gin.Context) {
	categories, err := c.usecase.Execute()
	if err != nil {
		// エラーレスポンスを生成する
		ErrorResponse(context, err)
		return
	}
	// CategoryエンティティからCategoryDTOに変換する
	results := []*dto.CategoryDTO{}
	for _, category := range categories {
		result, err := c.adapter.Convert(category)
		if err != nil {
			ErrorResponse(context, err)
			return
		}
		results = append(results, result)
	}
	// OKレスポンスを返す
	context.JSON(http.StatusOK, results)
}

// コンストラクタ
func NewcategoryListHandler(usecase usecase.CategoryList,
	adapter adapter.CategoryAdapter[*dto.CategoryDTO]) *categoryListHandler {
	return &categoryListHandler{
		usecase: usecase,
		adapter: adapter,
	}
}
