package usergin

import (
	"food-delivery-service/common"
	"food-delivery-service/component/hasher"
	userbiz "food-delivery-service/module/user/biz"
	usermodel "food-delivery-service/module/user/model"
	userstorage "food-delivery-service/module/user/storage"
	"food-delivery-service/plugin/tokenprovider"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := sc.MustGet(common.DBMain).(*gorm.DB)
		tokenProvider := sc.MustGet(common.JWTProvider).(tokenprovider.Provider)

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
