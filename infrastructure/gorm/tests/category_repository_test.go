package tests

import (
	"log"
	"service-exercise/infrastructure/errortype"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

// repository.CategoryRepositoryインターフェイスのテストドライバ
var _ = Describe("CategoryRepositoryインターフェイス実装のテスト\n", Ordered, func() {
	var env *InfrastructureTestEnvironment // テストに必要な環境
	BeforeAll(func() {                     // 前処理
		env = SetupTestEnvironment() // fxの起動と必要な環境の生成
	})
	AfterAll(func() { // 後処理
		TeardownTestEnvironment(env) // fxの停止
	})

	Describe("FindAll()メソッドを検証する\n", func() {
		When("商品カテゴリのスライスを取得する\n", func() {
			It("エラーはなく商品カテゴリのスライスが返される\n", func() {
				categories, err := env.category_repo.FindAll(env.db)
				// エラーがnilであることを検証する
				Expect(err).To(BeNil())
				// 取得件数が0件ではないことを検証する
				Expect(len(categories) > 0).To(Equal(true))
				for _, category := range categories {
					log.Println(category)
				}
			})
		})
		When("DBを停止した場合\n", func() {
			var db *gorm.DB
			BeforeEach(func() {
				db = env.db
				log.Println("データベース接続をクローズする")
				sqlDB, err := db.DB()
				if err != nil {
					panic("データベース接続の取得に失敗しました")
				}
				sqlDB.Close()
			})
			AfterEach(func() {
				env = SetupTestEnvironment() // fxの起動し直す
			})
			It("InternalErrorが返される\n", func() {
				departments, err := env.category_repo.FindAll(env.db)
				Expect(departments).To(BeNil())
				Expect(err).Should(BeAssignableToTypeOf(&errortype.InternalError{}))
			})
		})
	})
})
