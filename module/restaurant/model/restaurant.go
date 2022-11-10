package restaurantmodel

import (
	"food-delivery-service/common"
	"strings"
)

const EntityName = "Restaurant"

var (
	ErrNameCannotBeBlank = common.NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeBlank")
)

type Restaurant struct {
	common.SQLModel
	OwnerId     int                `json:"-" gorm:"column:owner_id;"`
	FakeOwnerId *common.UID        `json:"owner_id" gorm:"-"`
	Name        string             `json:"name" gorm:"column:name;"`
	Addr        string             `json:"address" gorm:"column:addr;"`
	LikedCount  int                `json:"liked_count" gorm:"column:liked_count;"`
	Liked       bool               `json:"liked" gorm:"-"`
	Owner       *common.SimpleUser `json:"owner" gorm:"foreignKey:OwnerId;PRELOAD:false;"`
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.SQLModel.Mask(common.DbTypeRestaurant)

	fakeOwnerId := common.NewUID(uint32(r.OwnerId), int(common.DbTypeUser), 1)
	r.FakeOwnerId = &fakeOwnerId

	if v := r.Owner; v != nil {
		v.Mask(common.DbTypeUser)
	}
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	OwnerId int    `json:"owner_id" gorm:"column:owner_id;"`
	Addr    string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Id = 0
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotBeBlank
	}

	return nil
}
