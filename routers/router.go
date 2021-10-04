package routers

import (
	"github.com/PenguinCats/unison-web-backend/docs"
	"github.com/PenguinCats/unison-web-backend/middleware/jwt"
	"github.com/PenguinCats/unison-web-backend/routers/api/auth"
	"github.com/PenguinCats/unison-web-backend/routers/api/message"
	"github.com/PenguinCats/unison-web-backend/routers/api/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api"

	apiG := r.Group("/api")

	apiG.POST("/auth/login_normal", auth.LoginNormal)

	apiG.POST("/user/profile", jwt.AuthLogin, user.GetUserProfile)

	apiG.POST("/message/message_inbox_profile_list", jwt.AuthLogin, message.GetMessageInboxProfileList)
	apiG.POST("/message/message_inbox_detail", jwt.AuthLogin, message.GetMessageInboxDetail)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
