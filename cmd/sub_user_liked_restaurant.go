package cmd

import (
	"context"
	"food-delivery-service/common"
	"food-delivery-service/component/asyncjob"
	restaurantstorage "food-delivery-service/module/restaurant/storage"
	"food-delivery-service/plugin/sdkgorm"
	"food-delivery-service/pubsub"
	appnats "food-delivery-service/pubsub/nats"
	"log"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var startSubUserLikedRestaurantCmd = &cobra.Command{
	Use:   "sub-user-liked-restaurant",
	Short: "Start a subscriber when user liked restaurant",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.DBMain)),
			goservice.WithInitRunnable(appnats.NewNATS(common.PluginNATS)),
			// goservice.WithInitRunnable(applocal.NewPubSub(common.PluginPubSub)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		ps := service.MustGet(common.PluginNATS).(pubsub.PubSub)
		ctx := context.Background()
		ch, _ := ps.Subscribe(ctx, common.TopicUserLikeRestaurant)

		for msg := range ch {
			db := service.MustGet(common.DBMain).(*gorm.DB)
			if restaurantId, ok := msg.Data()["restaurant_id"]; ok {
				job := asyncjob.NewJob(func(ctx context.Context) error {
					return restaurantstorage.NewSQLStore(db).IncreaseLikeCount(ctx, int(restaurantId.(float64)))
				})

				if err := asyncjob.NewGroup(true, job).Run(ctx); err != nil {
					log.Println(err)
				}
			}
		}
	},
}
