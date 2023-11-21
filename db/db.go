package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//新しいデータベースを作成しリターンする
func NewDB() *gorm.DB {
	//環境変数の読み込み時にエラーがあった際はプログラムの強制終了
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	//urlに環境変数を組み込んだpostgresURLを代入処理
	//postgresのurlは"postgres://user:pass@host:port/db"となる
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PW"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_DB"),
	)

	//postgresの開始とエラーの際の停止処理
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	//db開始したデータベースをリターン
	return db
}

//データベースのクローズ
func CloseDB(db *gorm.DB) {
	//db.DBは２変数返却する.
	//func (*gorm.DB).DB() (*sql.DB, error)としてポインター変数とerrorを出力するがerrorは使用しないためアンダーバーにて無効化
	sqlDB, _ := db.DB()

	//dbをクローズしてエラーがあった際にはプログラム強制終了
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}