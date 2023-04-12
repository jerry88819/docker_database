package model

import (
	"fmt"
	"time"
)

type MessageBoardData struct {
	UserName	string		`json:"userName"`
	Message		string		`json:"message"`
	Time		time.Time	`json:"time"`
} // ToolingType()

func ( MessageBoardData ) TableName() string {
	return "message_board_data"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

func CreateNewMessage( temp MessageBoardData ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateNewMessage()

func FindAllMessage() ( temp []MessageBoardData, err error ) {

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
		  tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return temp, err
	}

	if err := tx.Order( "time ASC").Find(&temp).Error; err != nil {
		tx.Rollback()
		return temp, err
	}

	fmt.Println("first", temp )

	if err := tx.Order( "time ASC").Find(&temp).Error; err != nil {
		tx.Rollback()
		return temp, err
	}

	fmt.Println("second", temp )

	return temp, tx.Commit().Error
} // FindAllTool()

func UpdateOneMessage( username string, temp MessageBoardData ) ( err error ) {
	fmt.Println(username)

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
		  tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&temp).Where("user_name = ?", username ).Updates(temp).Error; err != nil {
		tx.Rollback()
		return err
	}

	// time.Sleep(8 * time.Second)

	var ttt MessageBoardData

	if err := tx.Where("asd = ?", "" ).Find(&ttt).Error ; err != nil {
		fmt.Println("failed")
		tx.Rollback()
		return err
	}

	fmt.Println(ttt)

	return tx.Commit().Error
} // UpdateOneToolType()

func DeleteOneMessage( username string, inputdata string ) ( err error ) {
	temp := new( MessageBoardData )
    err = db.Where( "user_name = ? AND message = ?", username, inputdata ).Delete(&temp).Error
	return
} // DeleteOneMessage()
