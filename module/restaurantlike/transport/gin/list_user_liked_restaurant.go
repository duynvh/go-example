package restaurantlikegin

import (
	"food-delivery-service/common"
	restaurantlikebiz "food-delivery-service/module/restaurantlike/biz"
	restaurantlikemodel "food-delivery-service/module/restaurantlike/model"
	restaurantlikestore "food-delivery-service/module/restaurantlike/store"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ListUsers(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("restaurant_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		_ = paging.Validate()

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		store := restaurantlikestore.NewSQLStore(db)
		biz := restaurantlikebiz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(common.DbTypeUser)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
