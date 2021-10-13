package routers

import (
	"github.com/PenguinCats/unison-web-backend/middleware/jwt"
	"github.com/PenguinCats/unison-web-backend/routers/api/auth"
	"github.com/PenguinCats/unison-web-backend/routers/api/message"
	"github.com/PenguinCats/unison-web-backend/routers/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiG := r.Group("/api")

	apiG.POST("/auth/login_normal", auth.LoginNormal)

	apiG.POST("/user/profile", jwt.AuthLogin, user.GetUserProfile)
	apiG.POST("/user/search", jwt.AuthLogin, user.GetUserSearch)
	apiG.POST("/user/list", jwt.AuthAdmin, user.GetUserList)

	apiG.POST("/message/message_inbox_profile_list", jwt.AuthLogin, message.GetMessageInboxProfileList)
	apiG.POST("/message/message_inbox_detail", jwt.AuthLogin, message.GetMessageInboxDetail)
	apiG.POST("/message/add", jwt.AuthLogin, message.AddMessage)
	apiG.POST("/message/delete", jwt.AuthLogin, message.DeleteMessageAsUserView)
	apiG.POST("/message/mark_read", jwt.AuthLogin, message.MarkAsReadMessage)

	return r
}
