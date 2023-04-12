package model

type AccountInfo struct {
	Account		string		`json:"account"`
	Password	string		`json:"password"`
} // ToolingType()

func ( AccountInfo ) TableName() string {
	return "account_info"
} // TableName() 讓 gorm 知道確切的 Table 是哪一個

func CreateNewAccount( temp AccountInfo ) ( err error ) {
    err = db.Create(&temp).Error
	return
} // CreateNewAccount()

func CheckAdmin( temp AccountInfo ) ( err error ) {
    err = db.Where( "account = ? AND password = ?", temp.Account, temp.Password ).First(&temp).Error
	return
} // CheckAdmin() 