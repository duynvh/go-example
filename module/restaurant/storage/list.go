package restaurantstorage

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

func (store sqlStore) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	//offset := (paging.Page - 1) * paging.Limit

	var result []restaurantmodel.Restaurant

	db := store.db

	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id = ?", v)
	}

	if v := paging.FakeCursor; v != "" {
		if uid, err := common.FromBase58(v); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		}
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
		Count(&paging.Total).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		result[len(result)-1].Mask(true)
		paging.NextCursor = result[len(result)-1].FakeId.String()
	}

	return result, nil
}
