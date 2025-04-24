package tests

import (
	"log"
	"service-exercise/infrastructure/errortype"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// usecase.CategoryListインターフェイスのテストドライバ
var _ = Describe("CategoryListインターフェイス実装の品質検証\n", Ordered, func() {
	var env *ApplicationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	Context("商品カテゴリ一覧取得を検証する\n", func() {
		When("データベースが稼働している場合\n", func() {
			It("商品カテゴリが取得でき、エラーは無い\n", func() {
				categories, err := env.categtoryList.Execute()
				Expect(err).To(BeNil())
				Expect(len(categories) > 0).To(Equal(true))
				// 結果を出力する
				for _, category := range categories {
					log.Println(category.String())
				}
			})
		})
		When("データベースが停止している場合\n", func() {
			BeforeEach(func() {
				sqlDB, err := env.db.DB()
				if err != nil {
					panic("データベース接続の取得に失敗しました")
				}
				sqlDB.Close()
			})
			AfterEach(func() {
				env = SetupTestEnvironment()
			})
			It("商品カテゴリが取得できず、エラーが返される\n", func() {
				categories, err := env.categtoryList.Execute()
				Expect(categories).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err).Should(BeAssignableToTypeOf(&errortype.InternalError{}))
			})
		})
	})
})
