package routes

import (
	"log"
	"os"
	interfaces "servicev0/app/controller"
	"servicev0/src/infrastructure/config"
	"servicev0/src/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//API : Handler API Interfacing to FrontEnd
func API() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	conf := config.LoadAppConfig("postgres")
	dbdriver := conf.Driver
	host := conf.Host
	password := conf.Password
	user := conf.User
	dbname := conf.DBName
	port := conf.Port

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)

	if err != nil {
		panic(err)
	}

	users := interfaces.NewUsers(services.User)

	r := gin.Default()
	v1 := r.Group("/v1/api")
	//user routes
	v1.GET("/users", users.GetUsers)
	//Starting the application
	appPort := os.Getenv("PORT") //using heroku host
	if appPort == "" {
		appPort = "1123" //localhost
	}

	log.Fatal(r.Run(":" + appPort))
}
