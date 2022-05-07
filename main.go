package main

import (
	"ToDoList-Server/model"
	"ToDoList-Server/router"

	"github.com/labstack/echo/v4"
)

func main() {

	//DB接続とテーブルを作成
	sqlDB := model.DBConnection()

	defer sqlDB.Close()

	//echoのインスタンスを生成
	e := echo.New()

	//Routingを設定する関数　引数はecho.echo型であり、戻り値はerror型
	router.SetRouter(e)
}
