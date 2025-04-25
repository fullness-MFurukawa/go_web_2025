package main

import (
	"service-exercise/presentation"

	"go.uber.org/fx"
)

// @Title Go モダンWeb開発
// @Version 1.0
// @Description 商品と商品カテゴリを管理するAPIサービス
// @TermOfService http://localhost:8085/
// @Contact.name XXXX
// @Contact.url XXXX
// @Contact.email XXXX
// @Licence.name Apache 2.0
// @Licence.url https://www.apache.org/licenses/LICENSE-2.0
// @Host localhost:8085
// @Basepath /
func main() {
	app := fx.New(
		// プレゼンテーション層の依存関係を構築する
		presentation.PresentationModule,
	)
	app.Run()
}
