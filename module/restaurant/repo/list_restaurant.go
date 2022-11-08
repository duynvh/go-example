package restaurantrepo

import (
	"context"
	"food-delivery-service/common"
	restaurantmodel "food-delivery-service/module/restaurant/model"
)

type ListRestaurantStore interface {
	ListRestaurant(
		ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

func NewListRestaurantRepo(store ListRestaurantStore, uStore UserStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store, uStore: uStore}
}

type listRestaurantRepo struct {
	store  ListRestaurantStore
	uStore UserStore
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {
	result, err := repo.store.ListRestaurant(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	userIds := make([]int, len(result))

	for i := range result {
		userIds[i] = result[i].OwnerId
	}

	users, err := repo.uStore.GetUsers(ctx, userIds)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	// O(N^2)
	//for i := range result {
	//	for j := range users {
	//		if result[i].OwnerId == users[j].Id {
	//			result[i].Owner = &users[j]
	//			break
	//		}
	//	}
	//}

	// O(N)
	mapUser := make(map[int]*common.SimpleUser)

	for j, u := range users {
		mapUser[u.Id] = &users[j]
	}

	for i, item := range result {
		result[i].Owner = mapUser[item.OwnerId]
	}

	return result, nil
}
