package tests

import (
	"log"
	"service-exercise/infrastructure/errortype"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// usecase.ProductKeywordインターフェイスのテストドライバ
var _ = Describe("ProductKeywordインターフェイス実装の品質検証\n", Ordered, func() {
	var env *ApplicationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})

	Context("商品検索を検証する\n", func() {
		When("存在する商品キーワードを利用した場合\n", func() {
			It("該当する商品を取得でき、エラーは無い\n", func() {
				products, err := env.productKeyword.Execute("ボールペン")
				Expect(err).To(BeNil())
				Expect(len(products) > 0).To(Equal(true))
				// 結果を出力する
				for _, product := range products {
					log.Println(product.String())
				}
			})
		})
		When("存在しない商品キーワードを利用した場合\n", func() {
			It("商品を取得できず、エラーが返される\n", func() {
				products, err := env.productKeyword.Execute("山田")
				Expect(products).To(BeNil())
				Expect(err).NotTo(BeNil())
				Expect(err).Should(BeAssignableToTypeOf(&errortype.NotFoundError{}))
				Expect(err.Error()).To(Equal("キーワード:'山田'に該当する商品は見つかりませんでした。"))
			})
		})
	})
})
