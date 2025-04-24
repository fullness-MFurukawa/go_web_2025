package tests

import (
	"service-exercise/domain/adapter"
	"service-exercise/domain/repository"
	"service-exercise/infrastructure"
	"service-exercise/infrastructure/gorm/dbmodel"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
)

// Ginkgoによるテストスイート実行を開始する
func TestInfrastructureSuite(t *testing.T) {
	// テスト失敗時のハンドラをGinkgoのFail()関数に設定
	RegisterFailHandler(Fail)
	// テストスイート:gormパッケージのテスト"を実行
	RunSpecs(t, "インフラストラクチャ層:gormパッケージのテスト")
}

// インフラストラクチャ層のテストに必要な環境
type InfrastructureTestEnvironment struct {
	app           *fxtest.App
	db            *gorm.DB
	category_adpt adapter.CategoryAdapter[*dbmodel.CategoryModel]
	product_adpt  adapter.ProductAdapter[*dbmodel.ProductModel]
	category_repo repository.CategtoryRepository[*gorm.DB]
	product_repo  repository.ProductRepository[*gorm.DB]
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *InfrastructureTestEnvironment {
	var env InfrastructureTestEnvironment
	app := fxtest.New(
		GinkgoT(),
		infrastructure.InfrastructureModule,
		fx.Populate(&env.db),
		fx.Populate(&env.category_adpt),
		fx.Populate(&env.product_adpt),
		fx.Populate(&env.category_repo),
		fx.Populate(&env.product_repo),
	)
	env.app = app
	// fxコンテナを起動する
	env.app.RequireStart()
	return &env
}

// fxコンテナを停止する
func TeardownTestEnvironment(env *InfrastructureTestEnvironment) {
	env.app.RequireStop()
}
