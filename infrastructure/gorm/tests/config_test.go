package tests

import (
	"service-exercise/infrastructure/gorm/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewConfig()関数の検証\n", func() {
	When("有効なYAMLファイルの読取り\n", func() {
		It("Config構造体返される\n", func() {
			// NewConfig関数を呼び出し
			cfg, err := config.NewConfig()
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).NotTo(BeNil())
			Expect(cfg.DB.User).To(Equal("root"))
			Expect(cfg.DB.Password).To(Equal("password"))
			Expect(cfg.DB.Host).To(Equal("goweb_exercise_db"))
			Expect(cfg.DB.Port).To(Equal(3306))
			Expect(cfg.DB.DBName).To(Equal("exercisedb"))
			Expect(cfg.DB.Option).To(Equal("?charset=utf8mb4&parseTime=True&loc=Local"))
		})
	})
})
