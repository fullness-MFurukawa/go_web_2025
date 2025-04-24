package tests

import (
	"service-exercise/domain/model/categories"
	"service-exercise/infrastructure/gorm/dbmodel"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// adapter.CategoryAdapterインターフェイスのテストドライバ
var _ = Describe("CategoryAdapterインターフェイス実装のテスト\n", Ordered, func() {
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
				result, err := env.category_adpt.Convert(nil)
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("引数がnilのため、CategoryModelへの変換ができません。"))
			})
		})
		When("引数が有効な場合\n", func() {
			var category *categories.Category
			BeforeEach(func() {
				id, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				name, err := categories.NewCategoryName("文房具")
				Expect(err).NotTo(HaveOccurred())
				category, err = categories.NewCategory(id, name)
				Expect(err).NotTo(HaveOccurred())
			})
			It("正常に変換される\n", func() {
				result, err := env.category_adpt.Convert(category)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(&dbmodel.CategoryModel{
					ID:    0,
					ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name:  "文房具",
				}))
			})
		})
	})

	Describe("Restore()メソッドを検証する\n", func() {
		When("引数がnilの場合\n", func() {
			It("エラーを返す必要が返される\n", func() {
				result, err := env.category_adpt.Restore(nil)
				Expect(result).To(BeNil())
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("引数がnilのため、Categoryを復元できません。"))
			})
		})
		When("引数が有効な場合\n", func() {
			var (
				model    *dbmodel.CategoryModel
				category *categories.Category
			)
			BeforeEach(func() {
				model = &dbmodel.CategoryModel{
					ID:    0,
					ObjID: "b1524011-b6af-417e-8bf2-f449dd58b5c0",
					Name:  "文房具",
				}
				id, err := categories.NewCategoryId("b1524011-b6af-417e-8bf2-f449dd58b5c0")
				Expect(err).NotTo(HaveOccurred())
				name, err := categories.NewCategoryName("文房具")
				Expect(err).NotTo(HaveOccurred())
				category, err = categories.NewCategory(id, name)
				Expect(err).NotTo(HaveOccurred())
			})
			It("正常に変換される\n", func() {
				result, err := env.category_adpt.Restore(model)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(category))
			})
		})
	})
})
