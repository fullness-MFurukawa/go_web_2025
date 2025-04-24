package tests

import (
	"log"
	"service-exercise/application/errortype"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/model/products"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// usecase.ProductCreateインターフェイスのテストドライバ
var _ = Describe("ProductCreateインターフェイス実装の品質検証\n", Ordered, func() {
	var env *ApplicationTestEnvironment
	BeforeAll(func() {
		env = SetupTestEnvironment()
	})
	AfterAll(func() {
		TeardownTestEnvironment(env)
	})
	Context("商品登録を検証する\n", func() {
		When("存在しない商品を登録した場合\n", func() {
			var product *products.Product
			BeforeEach(func() {
				id := products.NewProductIdWithUUID()
				name, err := products.NewProductName("テスト商品")
				Expect(err).NotTo(HaveOccurred())
				price, err := products.NewProductPrice(200)
				Expect(err).NotTo(HaveOccurred())
				categoryId, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				category, err := categories.NewCategory(categoryId, nil)
				Expect(err).NotTo(HaveOccurred())
				product, err = products.NewProduct(id, name, price, category)
				Expect(err).NotTo(HaveOccurred())
			})
			It("商品が登録され、エラーは無い", func() {
				err := env.productRegister.Execute(product)
				Expect(err).To(BeNil())
				products, err := env.productKeyword.Execute("テスト商品")
				Expect(err).To(BeNil())
				// 結果を出力する
				for _, product := range products {
					log.Println(product.String())
				}
			})
		})
		When("存在する商品を登録した場合\n", func() {
			var product *products.Product
			BeforeEach(func() {
				id, err := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				Expect(err).NotTo(HaveOccurred())
				name, err := products.NewProductName("水性ボールペン(黒)")
				Expect(err).NotTo(HaveOccurred())
				price, err := products.NewProductPrice(120)
				Expect(err).NotTo(HaveOccurred())
				categoryId, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				category, err := categories.NewCategory(categoryId, nil)
				Expect(err).NotTo(HaveOccurred())
				product, err = products.NewProduct(id, name, price, category)
				Expect(err).NotTo(HaveOccurred())
			})
			It("商品は登録されず、エラーが返される\n", func() {
				err := env.productRegister.Execute(product)
				Expect(err).NotTo(BeNil())
				Expect(err).Should(BeAssignableToTypeOf(&errortype.ApplicationError{}))
				Expect(err.Error()).To(Equal("商品:水性ボールペン(黒)は、既に登録済です。"))
			})
		})
	})
})
