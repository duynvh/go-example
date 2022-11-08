package restaurantstorage

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store *sqlStore) InsertRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.PrepareForInsert()

	if err := store.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
