package config

import (
	"bytes"
	"embed"
	"fmt"
	"service-exercise/infrastructure/errortype"

	"gopkg.in/yaml.v3"
)

// データベース接続情報を表す構造体
type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DBName   string `yaml:"dbname"`
		Option   string `yaml:"option"`
	} `yaml:"db"`
}

//go:embed config.yml
var f embed.FS

// コンストラクタ
// config.ymlを読み込んで、Config構造体を生成する
func NewConfig() (*Config, error) {
	var cfg Config
	// YAMLファイルの読み込み
	yamlFile, err := f.ReadFile("config.yml")
	if err != nil {
		errortype.NewInternalError(
			fmt.Sprintf("config.yamlの読み込みに失敗しました: %v", err))
	}
	// YAMLファイルの値をConfig構造体のマッピングする
	if err := yaml.NewDecoder(bytes.NewReader(yamlFile)).Decode(&cfg); err != nil {
		errortype.NewInternalError(
			fmt.Sprintf("YAMLの解析に失敗しました: %v", err))
	}
	return &cfg, nil
}
