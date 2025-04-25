package handler

import (
	"net/http"
	"service-exercise/application/usecase"
	"service-exercise/domain/adapter"
	"service-exercise/presentation/dto"

	"github.com/gin-gonic/gin"
)

// 商品登録結果を返すリクエストハンドラ
type productRegisterHandler struct {
	usecase usecase.ProductRegister
	adapter adapter.ProductAdapter[*dto.ProductDTO]
}

// 商品登録をする
// @Tags 商品
// @Summary 商品を登録する
// @Description 新しい商品を登録する
// @Id register-product
// @Accept application/json
// @Produce json
// @Param body body dto.ProductDTO true "商品データ"
// @Success 200 {object} dto.ProductDTO
// @Failure 400 {object} errortype.DomainError
// @Failure 400 {object} errortype.ApplicationError
// @Failure 500 {object} errortype.InternalError
// @Router /product/register [post]
func (p *productRegisterHandler) Execute(context *gin.Context) {

	var dto dto.ProductDTO

	// JSON リクエストボディからProductDTOにバインド
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ProductDTOからProductエンティティを復元する
	product, err := p.adapter.Restore(&dto)
	if err != nil {
		// 復元に失敗:エラーレスポンスを返す
		ErrorResponse(context, err)
		return
	}
	// 従業員を登録する
	err = p.usecase.Execute(product)
	if err != nil {
		// 登録に失敗:エラーレスポンスを返す
		ErrorResponse(context, err)
		return
	}
	// 正常に登録されたのでEmployeeDTOを返す
	context.JSON(http.StatusOK, dto)
}

// コンストラクタ
func NewproductRegisterHandler(usecase usecase.ProductRegister,
	adapter adapter.ProductAdapter[*dto.ProductDTO]) *productRegisterHandler {
	return &productRegisterHandler{
		usecase: usecase,
		adapter: adapter,
	}
}
