package restaurantlikebiz

import (
	"context"
	"food-delivery-service/common"
	"food-delivery-service/component/asyncjob"
	restaurantlikemodel "food-delivery-service/module/restaurantlike/model"
	"github.com/200Lab-Education/go-sdk/logger"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	CheckUserLike(ctx context.Context, userId, restaurantId int) (bool, error)
}

type IncLikedCountResStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	//pb    pubsub.Pubsub
	incStore IncLikedCountResStore
}

func NewUserLikeRestaurantBiz(
	store UserLikeRestaurantStore,
	incStore IncLikedCountResStore,
	//	pb pubsub.Pubsub,
) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{
		store: store,
		//pb:    pb,
		incStore: incStore,
	}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {
	liked, err := biz.store.CheckUserLike(ctx, data.UserId, data.RestaurantId)

	if err != nil && err != common.ErrRecordNotFound {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	if liked {
		return restaurantlikemodel.ErrUserAlreadyLikedRestaurant(nil)
	}

	err = biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side effect
	go func() {
		defer common.Recover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
				logger.GetCurrent().GetLogger("user.like.restaurant").Errorln(err)
				return err
			}

			return nil
		}, asyncjob.WithName("IncreaseLikeCount"))

		if err := asyncjob.NewGroup(false, job).Run(ctx); err != nil {
			logger.GetCurrent().GetLogger("user.like.restaurant").Errorln(err)
		}
	}()

	//newMessage := pubsub.NewMessage(data)
	//biz.pb.Publish(ctx, common.TopicUserLikeRestaurant, newMessage)

	return nil
}
