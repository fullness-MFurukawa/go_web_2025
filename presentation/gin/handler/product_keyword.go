package handler

import (
	"net/http"
	"service-exercise/application/usecase"
	"service-exercise/domain/adapter"
	"service-exercise/presentation/dto"

	"github.com/gin-gonic/gin"
)

// 商品キーワード検索結果を返すリクエストハンドラ
type productKeywordHandler struct {
	usecase usecase.ProductKeyword
	adapter adapter.ProductAdapter[*dto.ProductDTO]
}

// キーワード検索した商品を取得する
// @Tags 商品
// @Summary 商品を取得する
// @Description キーワード検索した商品を取得する
// @Id keyword-product
// @Accept application/json
// @Produce json
// @Param keyword path string true "商品名(キーワード)"
// @Success 200 {object} []dto.ProductDTO
// @Failure 404 {object} errortype.NotFoundError
// @Failure 500 {object} errortype.InternalError
// @Router /product/keyword/{keyword} [GET]
func (e *productKeywordHandler) Execute(context *gin.Context) {
	keyword := context.Param("keyword") // パラメータを取得する
	products, err := e.usecase.Execute(keyword)
	if err != nil {
		ErrorResponse(context, err) // エラーレスポンスを返す
		return
	}
	results := []*dto.ProductDTO{}
	for _, product := range products {
		result, err := e.adapter.Convert(product)
		if err != nil {
			ErrorResponse(context, err) // 変換に失敗:エラーレスポンスを返す
			return
		}
		results = append(results, result)
	}
	context.JSON(http.StatusOK, results)
}

// コンストラクタ
func NewproductKeywordHandler(usecase usecase.ProductKeyword,
	adapter adapter.ProductAdapter[*dto.ProductDTO]) *productKeywordHandler {
	return &productKeywordHandler{
		usecase: usecase,
		adapter: adapter,
	}
}
