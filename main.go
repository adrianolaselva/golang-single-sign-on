package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "gopkg.in/go-oauth2/mysql.v3"
	"log"
	"oauth2/src/common"
	"oauth2/src/models"
	"oauth2/src/repository"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("file .env not found")
	}
	
	db := common.Database{}
	conn := *db.Connect()

	defer conn.Close()

	//hash := common.Hash{}
	dateCommon := common.DateCommon{}
	birthday := dateCommon.ConvertFromDateStr("1999-12-31")

	var user = &models.User{
		Uuid:      "8d42ee3e-5717-4b65-b0b6-218361f981b3",
		Username:  "adrianolaselva3",
		Email:     "adrianolaselva3@gmail.com",
		//Password:  hash.BCryptGenerate("123@mudar"),
		Birthday:  &birthday,
		Activated: true,
		DeletedAt: nil,
	}
	userRepository := repository.NewUserRepository(&conn)

	err = userRepository.Update(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}
