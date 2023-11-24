package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

//データベースを作成しmodelで作成した型のスキーマを作成する関数.
func main() {
	//dbを作成しインスタンス化する処理
	dbConn := db.NewDB()

	//プログラムが終了した際にマイグレーションが完了したことを表示する
	defer fmt.Println("Successfully Migrated")

	//データベースをクローズする処理
	defer db.CloseDB(dbConn)

	//データベースにUserとTaskスキーマを作成する.このファイルの一番大事な処理
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}