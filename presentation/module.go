package presentation

import (
	"log"
	"service-exercise/application"
	"service-exercise/presentation/gin/adapter"
	"service-exercise/presentation/gin/handler"
	"service-exercise/presentation/gin/preparation"

	"go.uber.org/fx"
)

// アプリケーション層の依存関係を構築する
var PresentationModule = fx.Options(
	application.ApplicationModule,
	fx.Provide(
		// Adapterの登録
		adapter.NewcategoryAdapaterGin,
		adapter.NewproductAdapterGin,
		// リクエストハンドラの登録
		handler.NewcategoryListHandler,
		handler.NewproductRegisterHandler,
		handler.NewproductKeywordHandler,
		handler.ProvideHandlers,
		// Router構造体の登録
		preparation.NewRouter,
	),
	// fx起動時、停止時の処理を登録
	fx.Invoke(preparation.RegisterHooks),
	fx.Invoke(setupEnd),
)

func setupEnd() {
	log.Println("プレゼンテーション層の構築が完了しました。")
}
