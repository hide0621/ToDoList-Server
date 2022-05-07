package router

import (
	"os"

	"github.com/labstack/echo/v4/middleware"

	_ "net/http"

	"github.com/labstack/echo/v4"
)

//Routingを設定する関数　引数はecho.echo型であり、戻り値はerror型
func SetRouter(e *echo.Echo) error {

	// 諸々の設定(*1)
	// APIが叩かれた時にログを出す
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
	// 予想外のエラーが発生した際でも、サーバーを落とさないようにする
	e.Use(middleware.Recover())
	// CORSに対応する
	e.Use(middleware.CORS())

	// APIを書く場所
	e.GET("/api/tasks", GetTasksHandler)
	e.POST("/api/tasks", AddTaskHandler)
	e.PUT("/api/tasks/:taskID", ChangeFinishedTaskHandler)
	e.DELETE("/api/tasks/:taskID", DeleteTaskHandler)

	// 8000番のポートを開く
	err := e.Start(":8000")
	return err

}
