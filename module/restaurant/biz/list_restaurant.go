package restaurantbiz

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type ListRestaurantRepo interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
	) ([]restaurantmodel.Restaurant, error)
}

func NewListRestaurantBiz(store ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

type listRestaurantBiz struct {
	store ListRestaurantRepo
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
