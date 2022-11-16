package cmd

import (
	"fmt"
	"food-delivery-service/cmd/handlers"
	"food-delivery-service/common"
	"food-delivery-service/middleware"
	userstorage "food-delivery-service/module/user/storage"
	"food-delivery-service/module/user/storage/grpcstore"
	"food-delivery-service/pubsub/localpb"

	appnats "food-delivery-service/pubsub/nats"

	// "food-delivery-service/plugin/sdkgorm"
	"food-delivery-service/plugin/appredis"
	appgrpc "food-delivery-service/plugin/remotecall/grpc"
	"food-delivery-service/plugin/remotecall/restful"
	"food-delivery-service/plugin/tokenprovider/jwt"
	user "food-delivery-service/proto"
	"net/http"
	"os"

	goservice "github.com/200Lab-Education/go-sdk"
	sdkgorm "github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("food-delivery"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
		goservice.WithInitRunnable(jwt.NewTokenJWTProvider(common.JWTProvider)),
		goservice.WithInitRunnable(restful.NewUserService()),
		goservice.WithInitRunnable(localpb.NewPubSub(common.PluginPubSub)),
		goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
		goservice.WithInitRunnable(appredis.NewRedisDB("redis", common.PluginRedis)),
		goservice.WithInitRunnable(appgrpc.NewGRPCServer(common.PluginGrpcServer)),
		goservice.WithInitRunnable(appgrpc.NewUserClient(common.PluginGrpcUserClient)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start an food delivery service",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()

		serviceLogger := service.Logger("service")
		service.MustGet(common.PluginGrpcServer).(interface {
			SetRegisterHdl(hdl func(*grpc.Server))
		}).SetRegisterHdl(func(server *grpc.Server) {
			dbConn := service.MustGet(common.DBMain).(*gorm.DB)
			user.RegisterUserServiceServer(server, grpcstore.NewGRPCStore(userstorage.NewSQLStore(dbConn)))
		})

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

	rootCmd.AddCommand(startSubUserLikedRestaurantCmd)
	rootCmd.AddCommand(startSubUserDislikedRestaurantCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
