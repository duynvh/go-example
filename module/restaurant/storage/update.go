package restaurantstorage

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"

	"gorm.io/gorm"
)

func (store *sqlStore) UpdateRestaurant(
	ctx context.Context,
	cond map[string]interface{},
	data *restaurantmodel.RestaurantUpdate,
) error {
	if err := store.db.Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (store *sqlStore) IncreaseLikeCount(ctx context.Context, id int) error {
	db := store.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (store *sqlStore) DecreaseLikeCount(ctx context.Context, id int) error {
	db := store.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).
		Update("liked_count", gorm.Expr("liked_count - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
