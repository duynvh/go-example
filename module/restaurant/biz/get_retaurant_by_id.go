package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type FindRestaurantStore interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
}

func NewFindRestaurantBiz(store FindRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

type findRestaurantBiz struct {
	store FindRestaurantStore
}

func (biz *findRestaurantBiz) FindRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return data, nil
}
