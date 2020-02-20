package main

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"oauth2/src/common"
	"oauth2/src/models"
	"oauth2/src/repository"
	"testing"
	"time"
)

var (
	conn *gorm.DB
	hash common.Hash
	dateCommon common.DateCommon
	userRepository repository.UserRepository
	birthday time.Time
	password string
	timeNowStr string
	idInserted string

	name      string
	lastName  string
	email     string
	username  string
)

func setup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("file .env not found")
	}

	db := common.Database{}
	conn = db.Connect()

	hash = common.NewHash()
	dateCommon = common.NewDateCommon()

	birthday = dateCommon.ConvertFromDateStr("1987-02-11")
	password, _ = hash.BCryptGenerate("123@mudar")
	timeNowStr = time.Now().Format("20060102150405")

	name      = "test_name_"+timeNowStr
	lastName  = "test_last_name_"+timeNowStr
	email     = "test_email_"+timeNowStr+"@test.com"
	username  = "test_username_"+timeNowStr

	userRepository = repository.NewUserRepository(conn)
}


func down() {
	defer conn.Close()
}

func TestCreateUser(t *testing.T) {
	setup()
	defer down()
	var roles []*models.Role
	roles = append(roles, &models.Role{
		ID: "1e57bdd1-9a1a-4264-9d3f-2aae2210b180",
		Name: "ADMINISTRADOR",
	})

	var user = &models.User{
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Username:  username,
		Password:  &password,
		Birthday:  &birthday,
		Activated: false,
		Roles:     roles,
	}
	err := userRepository.Create(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)
	idInserted = user.ID
}

func TestUpdateUser(t *testing.T) {
	setup()
	defer down()
	var roles []*models.Role
	roles = append(roles, &models.Role{ID: "1e57bdd1-9a1a-4264-9d3f-2aae2210b180"})
	user := &models.User{
		ID:		   idInserted,
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Username:  username,
		Password:  &password,
		Birthday:  &birthday,
		Activated: false,
	}

	err := userRepository.Update(user)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)
}

func TestFindByIdUser(t *testing.T) {
	setup()
	defer down()
	user, err := userRepository.FindById(idInserted)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)
}

func TestFindByUsernameUser(t *testing.T) {
	setup()
	defer down()
	user, err := userRepository.FindByUsername("adrianolaselva")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)
}

func TestFindByEmailUser(t *testing.T) {
	setup()
	defer down()
	user, err := userRepository.FindByEmail("adrianolaselva@gmail.com")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)
}

func TestFindByIdUserAndRoles(t *testing.T) {
	setup()
	defer down()
	user, err := userRepository.FindById("8d42ee3e-5717-4b65-b0b6-218361f981b3")
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(user.ID)

	for index, role := range user.Roles {
		log.Println(index, role.Name)
	}
}

//func createOAuthClient(t *testing.T) {
//	userRepository := repository.NewUserRepository(conn)
//	user, _ := userRepository.FindByUsername("adrianolaselva")
//	log.Println(user)
//}
