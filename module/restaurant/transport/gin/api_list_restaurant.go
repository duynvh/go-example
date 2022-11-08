package restaurantgin

import (
	"food-delivery-service/common"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantrepo "food-delivery-service/module/restaurant/repo"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	restaurantapi "food-delivery-service/module/restaurant/storage/remoteapi"
	"net/http"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		_ = paging.Validate()

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := restaurantstorage.NewSQLStore(db)
		//userStore := userstorage.NewSQLStore(db)
		userStore := restaurantapi.NewUserApi("http://localhost:3000")

		repo := restaurantrepo.NewListRestaurantRepo(store, userStore)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i := range result {
			result[i].Mask(true)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
