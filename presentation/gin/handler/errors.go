package handler

import (
	"net/http"
	appli "service-exercise/application/errortype"
	domain "service-exercise/domain/errortype"
	infra "service-exercise/infrastructure/errortype"

	"github.com/gin-gonic/gin"
)

// エラーレスポンス生成する
func ErrorResponse(ctx *gin.Context, err any) {
	switch e := err.(type) {
	case *domain.DomainError:
		ctx.JSON(http.StatusBadRequest, e.Error())
	case *appli.ApplicationError:
		ctx.JSON(http.StatusBadRequest, e.Error())
	case *infra.NotFoundError:
		ctx.JSON(http.StatusNotFound, e.Error())
	case *infra.InternalError:
		ctx.JSON(http.StatusInternalServerError, e.Error())
	default:
		ctx.JSON(http.StatusInternalServerError,
			"不明なエラーが検出されました。")
	}
}
