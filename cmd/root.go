package cmd

import (
	"fmt"
	"food-delivery-service/cmd/handlers"
	"food-delivery-service/common"
	"food-delivery-service/middleware"
	// "food-delivery-service/plugin/sdkgorm"
	"food-delivery-service/plugin/tokenprovider/jwt"
	"food-delivery-service/plugin/remotecall"
	goservice "github.com/200Lab-Education/go-sdk"
	sdkgorm "github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
		goservice.WithInitRunnable(remotecall.NewUserService()),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			engine.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": "pong"})
			})

			handlers.MainRoute(engine, service)
			handlers.InternalRoute(engine, service)
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	// TransAddPoint outenv as a sub command
	rootCmd.AddCommand(outEnvCmd)
	rootCmd.AddCommand(cronjob)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
