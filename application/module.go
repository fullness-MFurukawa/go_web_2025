package application

import (
	"log"
	"service-exercise/application/impls"
	"service-exercise/infrastructure"

	"go.uber.org/fx"
)

// アプリケーション層の依存関係を構築する
var ApplicationModule = fx.Options(
	// インフラストラクチャ層の依存定義
	infrastructure.InfrastructureModule,
	fx.Provide(
		// インターフェイス実装のコンストラクタ名を定義する
		impls.NewcategoryListImpl,
		impls.NewproductKeywordImpl,
		impls.NewproductRegisterImpl,
	),
	fx.Invoke(setupEnd),
)

func setupEnd() {
	log.Println("アプリケーション層の構築が完了しました。")
}
