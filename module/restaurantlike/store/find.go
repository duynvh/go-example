package restaurantlikestore

import (
	"context"
	"food-delivery-service/common"
	restaurantlikemodel "food-delivery-service/module/restaurantlike/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUserLike(ctx context.Context, userId, restaurantId int) (*restaurantlikemodel.Like, error) {
	var data restaurantlikemodel.Like

	if err := s.db.
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}

func (s *sqlStore) CheckUserLike(ctx context.Context, userId, restaurantId int) (bool, error) {
	var data restaurantlikemodel.Like

	if err := s.db.
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, common.ErrRecordNotFound
		}

		return false, common.ErrDB(err)
	}

	return true, nil
}
