package restaurantbiz

import (
	"context"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindRestaurant(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	DeleteRestaurant(
		ctx context.Context,
		cond map[string]interface{},
	) error
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func (biz *deleteRestaurantBiz) DeleteRestaurantById(ctx context.Context, id int) error {
	_, err := biz.store.FindRestaurant(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if err := biz.store.DeleteRestaurant(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}

	return nil
}
