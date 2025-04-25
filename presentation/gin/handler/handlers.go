package handler

// Handlers: Ginルーターに登録するハンドラーをグループ化
type Handlers struct {
	CategiryList    *categoryListHandler
	ProductRegister *productRegisterHandler
	ProductKeyword  *productKeywordHandler
}

// Handlers構造体のインスタンスをfxのコンテナに提供するための関数
func ProvideHandlers(
	categoryList *categoryListHandler,
	productRegister *productRegisterHandler,
	productKeyword *productKeywordHandler,
) Handlers {
	return Handlers{
		CategiryList:    categoryList,
		ProductRegister: productRegister,
		ProductKeyword:  productKeyword,
	}
}
