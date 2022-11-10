package restaurantlikegin

import (
	"food-delivery-service/common"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	restaurantlikebiz "food-delivery-service/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery-service/module/restaurantlike/model"
	restaurantlikestore "food-delivery-service/module/restaurantlike/store"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

// POST /v1/restaurants/:id/like // RPC-RestAPI
// POST /v1/restaurants/:id/liked-users // RestAPI

func UserLikeRestaurant(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)

		store := restaurantlikestore.NewSQLStore(db)
		incStore := restaurantstorage.NewSQLStore(db)
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, incStore)

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
