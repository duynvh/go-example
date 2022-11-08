package handlers

import (
	userinternal "food-delivery-service/module/user/transport/internalgin"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
)

func InternalRoute(router *gin.Engine, sc goservice.ServiceContext) {
	internal := router.Group("/internal")
	{
		internal.POST("/get-users-by-ids", userinternal.GetUserById(sc))
	}
}