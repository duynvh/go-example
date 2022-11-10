package restaurantlikegin

import (
	"food-delivery-service/common"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	restaurantlikebiz "food-delivery-service/module/restaurantlike/biz"
	restaurantlikestore "food-delivery-service/module/restaurantlike/store"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// DELETE /v1/restaurants/:id/dislike

func UserDislikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := sc.MustGet(common.DBMain).(*gorm.DB)

		store := restaurantlikestore.NewSQLStore(db)
		decStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUserDislikeRestaurantBiz(store, decStore)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
