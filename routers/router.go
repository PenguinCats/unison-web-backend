package routers

import (
	"github.com/PenguinCats/unison-web-backend/middleware/jwt"
	"github.com/PenguinCats/unison-web-backend/routers/api/auth"
	"github.com/PenguinCats/unison-web-backend/routers/api/host"
	"github.com/PenguinCats/unison-web-backend/routers/api/message"
	"github.com/PenguinCats/unison-web-backend/routers/api/permission"
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
	apiG.POST("/user/delete", jwt.AuthAdmin, user.DeleteUser)
	apiG.POST("/user/register", jwt.AuthAdmin, user.AddUser)
	apiG.POST("/user/edit", jwt.AuthAdmin, user.EditUser)
	apiG.POST("/user/pwd_edit", jwt.AuthAdmin, user.EditUserPassword)

	apiG.POST("/message/message_inbox_profile_list", jwt.AuthLogin, message.GetMessageInboxProfileList)
	apiG.POST("/message/message_inbox_detail", jwt.AuthLogin, message.GetMessageInboxDetail)
	apiG.POST("/message/add", jwt.AuthLogin, message.AddMessage)
	apiG.POST("/message/delete", jwt.AuthLogin, message.DeleteMessageAsUserView)
	apiG.POST("/message/mark_read", jwt.AuthLogin, message.MarkAsReadMessage)

	apiG.POST("/permission/list", jwt.AuthAdmin, permission.GetPermissionList)
	apiG.POST("/permission/user", jwt.AuthLogin, permission.GetPermissionGroupOfUser)
	apiG.POST("/permission/update_name", jwt.AuthAdmin, permission.UpdatePermissionGroupName)
	apiG.POST("/permission/delete", jwt.AuthAdmin, permission.DeletePermissionGroup)
	apiG.POST("/permission/add", jwt.AuthAdmin, permission.AddPermissionGroup)
	apiG.POST("/permission/host_bind", jwt.AuthAdmin, permission.BindPermissionGroupHost)
	apiG.POST("/permission/host", jwt.AuthAdmin, permission.GetHostOfPermissionGroup)

	apiG.POST("/host/all_list", jwt.AuthAdmin, host.GetHostAllList)
	apiG.POST("/host/update_ext", jwt.AuthAdmin, host.UpdateHostExt)
	apiG.POST("/host/update_add_token", jwt.AuthAdmin, host.UpdateHostAddToken)
	apiG.POST("/host/get_add_token", jwt.AuthAdmin, host.GetHostAddToken)

	return r
}
