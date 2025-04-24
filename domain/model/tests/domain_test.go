package tests

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Ginkgoによるテストスイート実行を開始する
func TestDomainSuite(t *testing.T) {
	// テスト失敗時のハンドラをGinkgoのFail関数に設定
	RegisterFailHandler(Fail)
	// テストスイート"ドメイン層:modelパッケージのテスト"を実行
	RunSpecs(t, "ドメイン層:modelパッケージのテスト")
}
