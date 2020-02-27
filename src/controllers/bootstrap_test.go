package controllers_test

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/repository"
	"oauth2/src/service"
	"os"
	"testing"
	"time"
)

// commons
var connection *gorm.DB
var hash common.Hash
var dateCommon common.DateCommon
var randomVal string
// repositories
var userRepository repository.UserRepository
var roleRepository repository.RoleRepository
// services
var userService service.UserService
var roleService service.RoleService
//controllers
var userController controllers.UserController
var healthController controllers.HealthController
var roleController controllers.RoleController

func TestMain(m *testing.M) {
	connection = NewConnection()
	userRepository = repository.NewUserRepository(connection)
	roleRepository = repository.NewRoleRepository(connection)
	userService = service.NewUserService(userRepository)
	roleService = service.NewRoleService(roleRepository)
	userController = controllers.NewUserController(userService)
	roleController = controllers.NewRoleController(roleService)
	randomVal = time.Now().Format("20060102150405")
	hash = common.NewHash()
	dateCommon = common.NewDateCommon()
	healthController = controllers.NewHealthController()
	defer connection.Close()
	os.Exit(m.Run())
}

func loadConfiguraton() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("file .env not found")
	}
}

func NewConnection() *gorm.DB {
	loadConfiguraton()
	db := common.Database{}
	conn := *db.Connect()
	return &conn
}
