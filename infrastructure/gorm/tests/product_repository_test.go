package tests

import (
	"log"
	"service-exercise/domain/model/categories"
	"service-exercise/domain/model/products"
	"service-exercise/infrastructure/errortype"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

// repository.ProductRepositoryインターフェイスのテストドライバ
var _ = Describe("ProductRepositoryインターフェイス実装のテスト\n", Ordered, func() {
	var env *InfrastructureTestEnvironment // テストに必要な環境
	BeforeAll(func() {                     // 前処理
		env = SetupTestEnvironment() // fxの起動と必要な環境の生成
	})
	AfterAll(func() { // 後処理
		TeardownTestEnvironment(env) // fxの停止
	})

	// Exists()メソッドのテスト
	Describe("Exists()メソッドを検証する\n", func() {
		When("存在する商品名を指定する\n", func() {
			var productName *products.ProductName
			// 存在する商品名を準備する
			BeforeEach(func() {
				if name, err := products.NewProductName("水性ボールペン(黒)"); err != nil {
					Expect(err).To(HaveOccurred())
				} else {
					productName = name
				}
			})
			It("trueとnilが返される\n", func() {
				result, err := env.product_repo.Exists(env.db, productName)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).To(BeNil())
				Expect(result).To(BeTrue())
			})
		})
		When("存在しない商品名を指定する\n", func() {
			var productName *products.ProductName
			// 存在する商品名を準備する
			BeforeEach(func() {
				if name, err := products.NewProductName("水性ボールペン"); err != nil {
					Expect(err).To(HaveOccurred())
				} else {
					productName = name
				}
			})
			It("falseとnilが返される\n", func() {
				result, err := env.product_repo.Exists(env.db, productName)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).To(BeNil())
				Expect(result).To(BeFalse())
			})
		})
	})

	// Create()メソッドのテスト
	Describe("Create()メソッドを検証する\n", func() {
		var db *gorm.DB
		BeforeAll(func() { // 前処理
			db = env.db.Begin() // トランザクションを開始する
		})
		AfterAll(func() { // 後処理
			db.Rollback() // トランザクションをロールバックする
		})

		When("存在しない商品Idで登録する\n", func() {
			var product *products.Product
			// 新しい商品を準備する
			BeforeEach(func() {
				productId := products.NewProductIdWithUUID()
				productName, err := products.NewProductName("水性ボールペン")
				Expect(err).NotTo(HaveOccurred())
				productPrice, err := products.NewProductPrice(200)
				Expect(err).NotTo(HaveOccurred())
				categoryId, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				category, err := categories.NewCategory(categoryId, nil)
				Expect(err).NotTo(HaveOccurred())
				product, err = products.NewProduct(productId, productName, productPrice, category)
				Expect(err).NotTo(HaveOccurred())
			})
			It("nilが返される\n", func() {
				err := env.product_repo.Create(db, product)
				Expect(err).To(BeNil())
			})
		})

		// FindByNameLike()メソッドのテスト
		Describe("FindByNameLike()メソッドを検証する\n", func() {
			When("存在する商品のキーワードを指定する", func() {
				It("該当する商品のスライスを返し、エラーはnilになる\n", func() {
					products, err := env.product_repo.FindByNameLike(env.db, "ボールペン")
					// errがnilであることを検証する
					Expect(err).To(BeNil())
					// 取得件数が0件ではないことを検証する
					Expect(len(products) > 0).To(Equal(true))
					for _, product := range products {
						log.Println(product)
					}
				})
			})
			When("存在しない商品のキーワードを指定する\n", func() {
				It("nilとerrtyps.NotFoundErrorが返される\n", func() {
					products, err := env.product_repo.FindByNameLike(env.db, "川")
					// errがnilであることを検証する
					Expect(products).To(BeNil())
					Expect(err).Should(BeAssignableToTypeOf(&errortype.NotFoundError{}))
					Expect(err.Error()).To(Equal(
						"キーワード:'川'に該当する商品は見つかりませんでした。"))
				})
			})
		})
	})
})
