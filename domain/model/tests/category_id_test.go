package tests

import (
	"service-exercise/domain/errortype"
	"service-exercise/domain/model/categories"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("値オブジェクト CategoryId \n", func() {
	Describe("NewCategoryIdWithUUID()関数を検証する \n", func() {
		It("有効なUUIDのCategoryIdが作成される \n", func() {
			categoryId := categories.NewCategoryIdWithUUID()
			// CategoryIdがnilでないことを検証する
			Expect(categoryId).NotTo(BeNil())
			// CategoryIdのvalueがUUID形式であることを検証する
			_, err := uuid.Parse(categoryId.Value())
			Expect(err).NotTo(HaveOccurred())
			// CategoryIdのvalueが36⽂字であることを検証する
			Expect(len(categoryId.Value())).To(Equal(36))
		})
	})

	Describe("NewCategoryId()関数を検証する \n", func() {
		When("有効なUUIDを使⽤した場合 \n", func() {
			It("エラーなくCategoryIdが返される \n", func() {
				uuid := "123e4567-e89b-12d3-a456-426614174000" // 仮の有効なUUID
				categoryId, err := categories.NewCategoryId(uuid)
				// エラーが発⽣していないことを検証する
				Expect(err).NotTo(HaveOccurred())
				// nilでないことを検証する
				Expect(categoryId).NotTo(BeNil())
				// valueフィールドのUUIDと仮のUUIDが等価であることを検証する
				Expect(categoryId.Value()).To(Equal(uuid))
			})
		})
		When("空文字の場合 \n", func() {
			It("エラーが返される \n", func() {
				categoryId, err := categories.NewCategoryId("")
				// エラーであることを検証する
				Expect(err).To(HaveOccurred())
				// エラーがerrortype.DomainError型であることを検証する
				Expect(err).To(BeAssignableToTypeOf(&errortype.DomainError{}))
				// エラーメッセージを検証する
				Expect(err.Error()).To(Equal("商品カテゴリIdは、空文字列であってはなりません。"))
				// categoryId変数がnilであることを検証する
				Expect(categoryId).To(BeNil())
			})
		})
	})

})
