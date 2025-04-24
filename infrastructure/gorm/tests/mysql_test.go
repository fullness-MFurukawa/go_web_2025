package tests

import (
	"fmt"
	"service-exercise/infrastructure"
	"service-exercise/infrastructure/gorm/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"gorm.io/gorm"
)

var _ = Describe("Connect()メソッドの検証\n", func() {
	When("有効な接続設定の場合\n", func() {
		It("データベースに正常に接続でき、*gorm.DBが返される", func() {
			cfg, _ := config.NewConfig()
			connector := config.NewMySQLConnector(cfg)
			db, err := connector.Connect()
			Expect(err).NotTo(HaveOccurred())
			Expect(db).NotTo(BeNil())
			dummyDB := &gorm.DB{}
			Expect(db).Should(BeAssignableToTypeOf(dummyDB))
		})
	})
	When("無効な接続設定の場合\n", func() {
		It("データベースへの接続に失敗し、エラーが返される", func() {
			// 接続設定を無効にする
			cfg, _ := config.NewConfig()
			cfg.DB.Port = -1
			connector := config.NewMySQLConnector(cfg)
			_, err := connector.Connect()
			fmt.Println(err.Error())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("内部の*sql.DBオブジェクトの取得に失敗しました: sql: database is closed"))
		})
	})
})

var _ = Describe("fxから*gorm.DBを取得\n", Ordered, func() {
	var db *gorm.DB
	var app *fxtest.App
	// 前処理
	BeforeEach(func() {
		app = fxtest.New(
			GinkgoT(),
			infrastructure.InfrastructureModule,
			fx.Populate(&db),
		)
		app.RequireStart()
	})
	// 後処理
	AfterEach(func() {
		app.RequireStop()
	})
	It("*gorm.DBが返される\n", func() {
		Expect(db).NotTo(BeNil())
		dummyDB := &gorm.DB{}
		Expect(db).Should(BeAssignableToTypeOf(dummyDB))
	})
})
