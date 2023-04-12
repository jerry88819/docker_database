package router

import (
	"net/http"
	"time"
	"github.com/jerry88819/docker_database/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	e *echo.Echo
)

func SetConfig() {

	e = echo.New()

	// set CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true, 
	}))

	e.Use(middleware.Logger()) 
	e.Use(middleware.Recover())
} // SetConfig()

func SetRouter() {

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})


	message := e.Group("/message")
	{
		message.GET("", handler.GetAllMessage )        // 取得所有留言
		message.PUT("", handler.PutMessage )       // 更新目前留言 其一 update
		message.POST("", handler.PostMessage )     // 建立新的留言內容 add
		message.DELETE("", handler.DeleteMessage ) // 刪除一筆留言 目前是用 名字 來搜尋需刪除的目標 delete
	}

	account := e.Group("/account")
	{
		account.GET("", handler.CheckAP ) 
		account.POST("", handler.PostAccount )   
	}


} // SerRouter()

func StartServer() {
	port := ":" + "8080"

	serverConfig := &http.Server{
		Addr:         port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	e.Logger.Fatal(e.StartServer(serverConfig))

	// e.Logger.Fatal(e.Start(":5000"))
} // StartServer()