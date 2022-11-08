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
	OwnerId int    `json:"owner_id" gorm:"column:owner_id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Addr    string `json:"address" gorm:"column:addr;"`
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.SQLModel.Mask(common.DbTypeRestaurant)
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
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
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
