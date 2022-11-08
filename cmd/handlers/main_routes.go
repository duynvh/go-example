package handlers

import (
	"food-delivery-service/common"
	"food-delivery-service/middleware"
	restaurantgin "food-delivery-service/module/restaurant/transport/gin"
	userstorage "food-delivery-service/module/user/storage"
	usergin "food-delivery-service/module/user/transport/gin"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func MainRoute(router *gin.Engine, sc goservice.ServiceContext) {
	dbConn := sc.MustGet(common.DBMain).(*gorm.DB)
	userStore := userstorage.NewSQLStore(dbConn)

	v1 := router.Group("/v1")
	{
		v1.GET("/admin",
			middleware.RequiredAuth(sc, userStore),
			middleware.RequiredRoles(sc, "admin", "mod"),
			func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"data": 1})
			})

		v1.POST("/register", usergin.Register(sc))
		v1.POST("/auth", usergin.Login(sc))
		v1.GET("/profile", middleware.RequiredAuth(sc, userStore), usergin.Profile(sc))

		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurantHandler(sc))
			restaurants.GET("", restaurantgin.ListRestaurant(sc))
			restaurants.GET("/:restaurant_id", restaurantgin.GetRestaurantHandler(sc))
			restaurants.PUT("/:restaurant_id", restaurantgin.UpdateRestaurantHandler(sc))
			restaurants.DELETE("/:restaurant_id", restaurantgin.DeleteRestaurantHandler(sc))
		}
	}
}
