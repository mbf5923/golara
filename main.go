package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"gorm-test/config"
	"gorm-test/routes"
	"gorm-test/utils"
	"gorm.io/gorm"
	"log"
)

func main() {
	router := setupRouter()
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))
}

func setupRouter() *gin.Engine {

	err := container.NamedSingleton("database", func() *gorm.DB {
		return config.DatabaseConnection()

	})
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{
		utils.GodotEnv("USER_GRPC_HOST"),
	})
	if utils.GodotEnv("GO_ENV") != "production" && utils.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if utils.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	routes.InitialRoutes(router)
	return router
}
