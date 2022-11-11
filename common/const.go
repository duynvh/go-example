package common

type DbType int

const (
	DbTypeRestaurant DbType = 1
	DbTypeUser       DbType = 2
)

const (
	CurrentUser = "user"

	DBMain            = "mysql"
	JWTProvider       = "jwt"
	PluginUserService = "user-service"
	PluginPubSub = "pubsub"
	PluginNATS = "nats"

	// PubSub Topics
	TopicUserLikeRestaurant = "restaurant.liked"
	TopicUserDislikeRestaurant = "restaurant.disliked"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}

func (p TokenPayload) Role() string {
	return p.URole
}
