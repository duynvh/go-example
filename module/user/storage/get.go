package userstorage

import (
	"context"
	"food-delivery-service/common"
)

func (s *sqlStore) GetUsers(ctx context.Context, ids []int) ([]common.SimpleUser, error) {
	var result []common.SimpleUser

	if err := s.db.Table(common.SimpleUser{}.TableName()).
		Where("id in (?)", ids).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
