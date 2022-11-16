package grpcstore

import (
	"context"
	"food-delivery-service/common"
	user "food-delivery-service/proto"
)

type UserStore interface {
	GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error)
}

type grpcStore struct {
	dbStore UserStore
}

func NewGRPCStore(dbStore UserStore) *grpcStore {
	return &grpcStore{dbStore: dbStore}
}

func (s *grpcStore) GetUserByIds(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	userIds := make([]int, len(request.GetUserIds()))

	for i := range userIds {
		userIds[i] = int(request.GetUserIds()[i])
	}

	rs, err := s.dbStore.GetUsers(ctx, userIds)

	if err != nil {
		return nil, err
	}

	users := make([]*user.User, len(rs))

	for i, item := range rs {
		item.Mask(common.DbTypeUser)

		users[i] = &user.User{
			Id:        item.FakeId.String(),
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Role:      item.Role,
		}
	}

	return &user.UserResponse{Users: users}, nil
}
