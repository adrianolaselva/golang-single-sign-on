package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/go-oauth2/mysql.v3"
	"oauth2/src"
)

func main() {
	app := src.Bootstrap{}
	app.Run()
}

//func main() {
//	//err := godotenv.Load(".env")
//	//if err != nil {
//	//	log.Println("file .env not found")
//	//}
//	//
//	//db := common.Database{}
//	//conn := *db.Connect()
//	//
//	//defer conn.Close()
//	//
//	//hash := common.Hash{}
//	//dateCommon := common.DateCommon{}
//	//birthday := dateCommon.ConvertFromDateStr("1987-02-11")
//	//password := hash.BCryptGenerate("123@mudar")
//	//
//	//var user = &models.User{
//	//	ID:      	"8d42ee3e-5717-4b65-b0b6-218361f981b3",
//	//	Name: 		"Adriano",
//	//	LastName: 	"Moreira La Selva",
//	//	Username:  	"adrianolaselva",
//	//	Email:     	"adrianolaselva@gmail.com",
//	//	Password:  	&password,
//	//	Birthday:  	&birthday,
//	//	Activated: 	true,
//	//	ExpiresAt: 	nil,
//	//	DeletedAt: 	nil,
//	//}
//	//
//	//userRepository := repository.NewUserRepository(&conn)
//	//
//	//err = userRepository.Update(user)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//
//	//fmt.Println(user)
//}
