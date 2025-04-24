package tests

import (
	"service-exercise/domain/model/categories"
	"service-exercise/domain/model/products"
	"service-exercise/infrastructure/gorm/dbmodel"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// adapter.ProductAdapterインターフェイスのテストドライバ
var _ = Describe("ProductAdapterインターフェイス実装のテスト\n", Ordered, func() {
	var env *InfrastructureTestEnvironment // テストに必要な環境
	BeforeAll(func() {                     // 前処理
		env = SetupTestEnvironment() // fxの起動と必要な環境の生成
	})
	AfterAll(func() { // 後処理
		TeardownTestEnvironment(env) // fxの停止
	})

	Describe("Convert()メソッドを検証する\n", func() {
		When("引数がnilの場合\n", func() {
			It("エラーを返す必要が返される\n", func() {
				result, err := env.product_adpt.Convert(nil)
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("引数がnilのため、ProductModelに変換できません。"))
			})
		})
		When("引数が有効な場合\n", func() {
			var product *products.Product
			BeforeEach(func() {
				id, err := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				Expect(err).NotTo(HaveOccurred())
				name, err := products.NewProductName("ボールペン")
				Expect(err).NotTo(HaveOccurred())
				price, err := products.NewProductPrice(200)
				Expect(err).NotTo(HaveOccurred())
				cid, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				cname, err := categories.NewCategoryName("文房具")
				Expect(err).NotTo(HaveOccurred())
				category, err := categories.NewCategory(cid, cname)
				Expect(err).NotTo(HaveOccurred())
				product, err = products.NewProduct(id, name, price, category)
				Expect(err).NotTo(HaveOccurred())
			})
			It("正常に変換される\n", func() {
				result, err := env.product_adpt.Convert(product)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(&dbmodel.ProductModel{
					ID:         0,
					ObjID:      "ac413f22-0cf1-490a-9635-7e9ca810e544",
					Name:       "ボールペン",
					Price:      200,
					CategoryID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
				}))
			})
		})
	})
	Describe("Restore()メソッドを検証する\n", func() {
		When("引数がnilの場合\n", func() {
			It("エラーを返す必要が返される\n", func() {
				result, err := env.product_adpt.Restore(nil)
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("引数がnilのため、Productを復元できません。"))
			})
		})
		When("引数が有効な場合\n", func() {
			var (
				model   *dbmodel.ProductModel
				product *products.Product
			)
			BeforeEach(func() {
				model = &dbmodel.ProductModel{
					ID:    0,
					ObjID: "ac413f22-0cf1-490a-9635-7e9ca810e544",
					Name:  "ボールペン",
					Price: 200,
					Category: dbmodel.CategoryModel{
						ID:    0,
						ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
						Name:  "文房具",
					},
				}
				id, err := products.NewProductId("ac413f22-0cf1-490a-9635-7e9ca810e544")
				Expect(err).NotTo(HaveOccurred())
				name, err := products.NewProductName("ボールペン")
				Expect(err).NotTo(HaveOccurred())
				price, err := products.NewProductPrice(200)
				Expect(err).NotTo(HaveOccurred())
				cid, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				cname, err := categories.NewCategoryName("文房具")
				Expect(err).NotTo(HaveOccurred())
				category, err := categories.NewCategory(cid, cname)
				Expect(err).NotTo(HaveOccurred())
				product, err = products.NewProduct(id, name, price, category)
				Expect(err).NotTo(HaveOccurred())
			})
			It("正常に変換される\n", func() {
				result, err := env.product_adpt.Restore(model)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(product))
			})
		})
	})
})
