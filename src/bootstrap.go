package src

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/middlewares"
	"oauth2/src/repository"
	"oauth2/src/routes"
	"oauth2/src/service"
	"oauth2/src/service/oauth"
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
	signature := os.Getenv("SSO_JWT_SIGNATURE")

	router := mux.NewRouter().StrictSlash(true)

	// Initialize repositories
	userRepository := repository.NewUserRepository(&conn)
	roleRepository := repository.NewRoleRepository(&conn)
	clientRepository := repository.NewClientRepository(&conn)
	authCodeRepository := repository.NewAuthCodeRepository(&conn)
	refreshTokenRepository := repository.NewRefreshTokenRepository(&conn)
	accessTokenRepository := repository.NewAccessTokenRepository(&conn)

	// Inicialize services
	userService := service.NewUserService(userRepository)
	roleService := service.NewRoleService(roleRepository)

	// OAuth2 implementation
	authFlow := oauth.NewAuthFlow(
			jwt.SigningMethodHS256,
			signature,
			userRepository,
			clientRepository,
			authCodeRepository,
			refreshTokenRepository,
			accessTokenRepository,
		)

	// Initialize controllers
	healthController := controllers.NewHealthController()
	swaggerController := controllers.NewSwaggerController()
	oauthController := controllers.NewOAuthController(authFlow)
	userController := controllers.NewUserController(userService)
	roleController := controllers.NewRoleController(roleService)

	// Initialize middlewares
	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(userService, authFlow)

	routeCommon := common.Route{}
	router.PathPrefix("/app").Handler(http.StripPrefix("/app", http.FileServer(http.Dir("./public/dist/app"))))
	routeCommon.Initialize(router, routes.NewHealthRoutes(healthController).Routes())
	routeCommon.Initialize(router, routes.NewSwaggerRoutes(swaggerController).Routes())
	routeCommon.Initialize(router, routes.NewOAuthRoutes(oauthController, authenticationMiddleware).Routes())
	routeCommon.Initialize(router, routes.NewUserRoutes(userController, authenticationMiddleware).Routes())
	routeCommon.Initialize(router, routes.NewRoleRoutes(roleController, authenticationMiddleware).Routes())

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/docs.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("list"),
		httpSwagger.DomID("#swagger-ui"),
	))

	headers := handlers.AllowedHeaders([]string{
		"Content-Type",
		"X-Request",
		"Location",
		"Entity",
		"Accept",
	})
	methods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
	})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Printf(fmt.Sprintf("Server started on port 0.0.0.0:%s", appPort))

	err = http.ListenAndServe(fmt.Sprintf(":%s", appPort), handlers.CORS(headers, methods, origins)(router))

	if err != nil {
		log.Fatal(err)
	}
}