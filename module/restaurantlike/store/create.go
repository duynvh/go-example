package restaurantlikestore

import (
	"context"
	"food-delivery-service/common"
	restaurantlikemodel "food-delivery-service/module/restaurantlike/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantlikemodel.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	//db.Exec("Update restaurants SET liked_count = liked_count + 1 where id = ?", data.RestaurantId)

	return nil
}
