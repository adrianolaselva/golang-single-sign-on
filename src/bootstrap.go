package src

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/middlewares"
	"oauth2/src/repository"
	"oauth2/src/routes"
	"oauth2/src/service"
	"os"
	"time"
)

type Bootstrap struct {

}

func (a * Bootstrap) Run() {
	_, _ = time.LoadLocation("America/Sao_Paulo")

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("file .env not found")
	}

	db := common.Database{}
	conn := *db.Connect()

	defer conn.Close()

	appPort := os.Getenv("SSO_PORT")

	router := mux.NewRouter().StrictSlash(true)

	// Initialize repositories
	userRepository := repository.NewUserRepository(&conn)
	roleRepository := repository.NewRoleRepository(&conn)

	// Inicialize services
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)

	// Initialize controllers
	healthController := controllers.NewHealthController()
	oauthController := controllers.NewOAuthController()
	userController := controllers.NewUserController(userService)
	roleController := controllers.NewRoleController(roleService)

	// Initialize middlewares
	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(userService)

	routeCommon := common.Route{}
	routeCommon.Initialize(router, routes.NewHealthRoutes(healthController).Routes())
	routeCommon.Initialize(router, routes.NewOAuthRoutes(oauthController, authenticationMiddleware).Routes())
	routeCommon.Initialize(router, routes.NewUserRoutes(userController, authenticationMiddleware).Routes())
	routeCommon.Initialize(router, routes.NewRoleRoutes(roleController, authenticationMiddleware).Routes())

	headers := handlers.AllowedHeaders([]string{"Content-Type", "X-Request", "Location", "Entity", "Accept"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Printf(fmt.Sprintf("Server started on port 0.0.0.0:%s", appPort))

	err = http.ListenAndServe(fmt.Sprintf(":%s", appPort), handlers.CORS(headers, methods, origins)(router))

	if err != nil {
		log.Fatal(err)
	}
}