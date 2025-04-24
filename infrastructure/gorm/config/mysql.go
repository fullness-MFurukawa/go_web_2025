package config

import (
	"fmt"
	"service-exercise/infrastructure/connector"
	"service-exercise/infrastructure/errortype"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLデータベース接続構造体
type MySQLConnector struct {
	Config *Config
}

// DatabaseConnectorインターフェイスの実装
// データベース接続
func (m *MySQLConnector) Connect() (*gorm.DB, error) {
	// データベース接続を生成する
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		m.Config.DB.User, m.Config.DB.Password, m.Config.DB.Host,
		m.Config.DB.Port, m.Config.DB.DBName, m.Config.DB.Option)

	// データベース接続(MySQLドライバのOpen関数)
	dialector := mysql.Open(dsn) // データベースに接続する
	// データベース接続からコネクションプールを生成する
	// GORMのデータベース接続
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		errortype.NewInternalError(
			fmt.Sprintf("データベースへの接続に失敗しました: %v", err))
	}

	// *sql.DB(コネクションプール)を取得する
	// 正しく接続できたか?
	sqlDB, err := db.DB()
	if err != nil {
		errortype.NewInternalError(
			fmt.Sprintf("内部の*sql.DBオブジェクトの取得に失敗しました: %v",
				err))
	}

	// コネクションプールの設定をカスタマイズする
	sqlDB.SetMaxIdleConns(10)           // アイドル状態の接続の最大数
	sqlDB.SetMaxOpenConns(100)          // オープン可能な接続の最大数
	sqlDB.SetConnMaxLifetime(time.Hour) // 接続の最大ライフタイム
	// 生成されたSQLをログ出力する
	db.Logger = db.Logger.LogMode(logger.Info)
	// データベース接続の確認
	if err := sqlDB.Ping(); err != nil {
		return nil,
			errortype.NewInternalError(
				fmt.Sprintf("内部の*sql.DBオブジェクトの取得に失敗しました: %v",
					err))
	}
	return db, nil
}

// コンストラクタ
// 戻り値: connector.DatabaseConnectorのジェネリクスに*gorm.DBを指定
func NewMySQLConnector(cfg *Config) connector.DatabaseConnector[*gorm.DB] {
	return &MySQLConnector{Config: cfg}
}
