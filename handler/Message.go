package handler

import (
	"fmt"
	"net/http"
	"time"
	"github.com/jerry88819/docker_database/model"
	"github.com/labstack/echo"
)

func PostMessage( c echo.Context ) error {
	temp := new( model.MessageBoardData )

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.Time = time.Now()

	fmt.Printf( temp.Message )
	fmt.Printf( temp.UserName )
	fmt.Println( temp.Time )

	err := model.CreateNewMessage( *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create Message Fail !!!")
	} // if()

	return c.JSON(http.StatusOK, "Success" )
} // postToolType()

func GetAllMessage(c echo.Context) error {

	ans, err := model.FindAllMessage()
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Cannot get message information from database")
	} // if()

	return c.JSON(http.StatusOK, ans)
} // GetAllMessage()

func PutMessage(c echo.Context) error {
	temp := new( model.MessageBoardData )
	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	temp.Time = time.Now()

    err := model.UpdateOneMessage( temp.UserName, *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Update Message Fail !!!")
	} // if()

	return c.JSON(http.StatusOK, "Success")
} // PutMessage()

func DeleteMessage(c echo.Context) error {
	temp := new( model.MessageBoardData )
	if err := c.Bind(temp); err != nil { // 跟前端過來的 json 檔案 bind 住
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()
	
	err := model.DeleteOneMessage( temp.UserName, temp.Message )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String( http.StatusInternalServerError, "Delete Message Fail !!!")
	} // if()

	return c.JSON( http.StatusOK, "Delete Success" )
} // DeleteMessage()
