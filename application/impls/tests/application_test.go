package tests

import (
	"service-exercise/application"
	"service-exercise/application/usecase"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
)

func TestApplicationSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "アプリケーション層のテストスイート")
}

type ApplicationTestEnvironment struct {
	app             *fxtest.App
	db              *gorm.DB
	categtoryList   usecase.CategoryList
	productKeyword  usecase.ProductKeyword
	productRegister usecase.ProductRegister
}

// テストに必要な環境を準備する
func SetupTestEnvironment() *ApplicationTestEnvironment {
	var env ApplicationTestEnvironment
	app := fxtest.New(
		GinkgoT(),
		application.ApplicationModule,
		fx.Populate(&env.db),
		fx.Populate(&env.categtoryList),
		fx.Populate(&env.productKeyword),
		fx.Populate(&env.productRegister),
	)
	env.app = app
	// fxコンテナを起動する
	env.app.RequireStart()
	return &env
}

// fxコンテナを停止する
func TeardownTestEnvironment(env *ApplicationTestEnvironment) {
	env.app.RequireStop()
}
