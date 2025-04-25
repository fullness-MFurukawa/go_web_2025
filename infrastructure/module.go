package infrastructure

import (
	"log"
	"service-exercise/infrastructure/gorm/config"
	"service-exercise/infrastructure/gorm/repimpl"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// データベース接続の*gorm.DBをfxに保持する
// DBModule:fx.Option
var DBModule = fx.Provide(func() (*gorm.DB, error) {
	// データベース接続情報を取得する
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	// MySQLに接続するコネクタを取得する
	connector := config.NewMySQLConnector(conf)
	// データベース接続した結果の*gorm.DBを
	return connector.Connect()
})

// インフラストラクチャ層の依存関係を構築する
var InfrastructureModule = fx.Options(
	DBModule, // データベース接続の*gorm.DBをfxに保持する
	fx.Provide(
		/*** 演習-15で追加する ***/
		// AdapterとRepositoryを生成して登録する
		repimpl.NewcategoryAdapterImpl,
		repimpl.NewproductAdapterImpl,
		repimpl.NewcategtoryRepositoryImpl,
		//repimpl.NewcategoryRepositoryImpl2,
		repimpl.NewproductRepositoryImpl,
	), // 内部的なインスタンス生成
	fx.Invoke(setupEnd), // 依存関係構築完了メッセージを出力する
)

func setupEnd() {
	log.Println("インフラストラクチャ層の構築が完了しました。")
}
