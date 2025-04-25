package preparation

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

// fxコンテナのライフサイクル
func RegisterHooks(lifecycle fx.Lifecycle, router *Router) {
	lifecycle.Append(
		fx.Hook{
			// fxコンテナ起動時の処理
			OnStart: func(context.Context) error {
				fmt.Println("演習APIの開始 Port:8085 !!!")
				go router.Engine.Run(":8085") //    Ginの起動
				return nil
			},
			// fxコンテナ停止時の処理
			OnStop: func(context.Context) error {
				fmt.Println("演習API停止 !!!")
				return nil
			},
		},
	)
}
