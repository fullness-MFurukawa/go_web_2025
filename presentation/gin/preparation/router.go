package preparation

import (
	_ "service-exercise/docs" // -- 追加 --
	"service-exercise/presentation/gin/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	sfiles "github.com/swaggo/files"      // -- 追加 ---
	gswag "github.com/swaggo/gin-swagger" // -- 追加 --
)

// ルーティング構造体
type Router struct {
	Engine *gin.Engine
}

// コンストラクタ
// handler.Handlers:リクエストハンドラをグループ化した構造体
func NewRouter(handlers handler.Handlers) *Router {

	gin.SetMode(gin.ReleaseMode) // リリースモードに設定する
	engine := gin.Default()      // デフォルトエンジンを生成する

	// CORSミドルウェアの設定（全オリジン許可、全メソッド許可）
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // すべてのオリジンからのアクセスを許可
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * 60 * 60, // 12時間キャッシュ
	}))

	// -- 追加 ---
	// SwaggerUIのエンドポイントとハンドラの登録
	url := gswag.URL("http://localhost:8085/swagger/doc.json")
	engine.GET("/swagger/*any", gswag.WrapHandler(sfiles.Handler, url))

	// エンドポイントとリクエストハンドラの登録
	engine.GET("/category/list", handlers.CategiryList.Execute)
	engine.GET("/product/keyword/:keyword", handlers.ProductKeyword.Execute)
	engine.POST("/product/register", handlers.ProductRegister.Execute)
	router := Router{Engine: engine}
	return &router
}
