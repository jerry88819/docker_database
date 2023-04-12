package handler

import (
	"fmt"
	"net/http"
	"github.com/jerry88819/docker_database/model"
	"github.com/labstack/echo"
)

func CheckAP(c echo.Context) error {
    temp := new( model.AccountInfo )

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	fmt.Println( temp.Account )
	fmt.Println( temp.Password )

	err := model.CheckAdmin( *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		// return c.String(http.StatusInternalServerError, "Cannot get message information from database")
		return c.String(http.StatusOK, "Wrong account!" )
	} // if()

	return c.String(http.StatusOK, "ADMIN CONFIRM")
} // GetAllMessage()

func PostAccount( c echo.Context ) error {
	temp := new( model.AccountInfo )

	if err := c.Bind(temp); err != nil {
		fmt.Printf("error =>%v\n", err)
		return err
	} // if()

	fmt.Printf( temp.Account )
	fmt.Printf( temp.Password )

	err := model.CreateNewAccount( *temp )
	if err != nil {
		fmt.Printf("error =>%v\n", err)
		return c.String(http.StatusInternalServerError, "Create Account Fail !!!")
	} // if()

	return c.JSON(http.StatusOK, "Success" )
} // PostAccount()