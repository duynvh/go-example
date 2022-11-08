package restaurantgin

import (
	"context"
	"food-delivery-service/common"
	restaurantbiz "food-delivery-service/module/restaurant/biz"
	restaurantmodel "food-delivery-service/module/restaurant/model"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type mockCreateStore struct{}

func (mockCreateStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.Id = 20
	return nil
}

func CreateRestaurantHandler(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		storage := restaurantstorage.NewSQLStore(db)
		//storage := &mockCreateStore{}
		biz := restaurantbiz.NewCreateRestaurantBiz(storage)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		data.Mask(common.DbTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
