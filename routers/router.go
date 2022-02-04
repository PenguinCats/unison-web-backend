package routers

import (
	"github.com/PenguinCats/unison-web-backend/middleware/jwt"
	"github.com/PenguinCats/unison-web-backend/routers/api/auth"
	"github.com/PenguinCats/unison-web-backend/routers/api/container"
	"github.com/PenguinCats/unison-web-backend/routers/api/host"
	"github.com/PenguinCats/unison-web-backend/routers/api/message"
	"github.com/PenguinCats/unison-web-backend/routers/api/operation"
	"github.com/PenguinCats/unison-web-backend/routers/api/permission"
	"github.com/PenguinCats/unison-web-backend/routers/api/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiG := r.Group("/api")

	apiG.POST("/auth/login_normal", auth.LoginNormal)

	apiUser := apiG.Group("/user")
	apiUser.POST("/profile", jwt.AuthLogin, user.GetUserProfile)
	apiUser.POST("/search", jwt.AuthLogin, user.GetUserSearch)
	apiUser.POST("/list", jwt.AuthAdmin, user.GetUserList)
	apiUser.POST("/delete", jwt.AuthAdmin, user.DeleteUser)
	apiUser.POST("/register", jwt.AuthAdmin, user.AddUser)
	apiUser.POST("/edit", jwt.AuthAdmin, user.EditUser)
	apiUser.POST("/pwd_edit", jwt.AuthAdmin, user.EditUserPassword)

	apiMessage := apiG.Group("/message")
	apiMessage.POST("/message_inbox_profile_list", jwt.AuthLogin, message.GetMessageInboxProfileList)
	apiMessage.POST("/message_inbox_detail", jwt.AuthLogin, message.GetMessageInboxDetail)
	apiMessage.POST("/add", jwt.AuthLogin, message.AddMessage)
	apiMessage.POST("/delete", jwt.AuthLogin, message.DeleteMessageAsUserView)
	apiMessage.POST("/mark_read", jwt.AuthLogin, message.MarkAsReadMessage)

	apiPermission := apiG.Group("/permission")
	apiPermission.POST("/list", jwt.AuthAdmin, permission.GetPermissionList)
	apiPermission.POST("/user", jwt.AuthLogin, permission.GetPermissionGroupOfUser)
	apiPermission.POST("/update_name", jwt.AuthAdmin, permission.UpdatePermissionGroupName)
	apiPermission.POST("/delete", jwt.AuthAdmin, permission.DeletePermissionGroup)
	apiPermission.POST("/add", jwt.AuthAdmin, permission.AddPermissionGroup)
	apiPermission.POST("/host_bind", jwt.AuthAdmin, permission.BindPermissionGroupHost)
	apiPermission.POST("/host", jwt.AuthAdmin, permission.GetHostOfPermissionGroup)

	apiHost := apiG.Group("/host")
	apiHost.POST("/all_list", jwt.AuthAdmin, host.GetHostAllList)
	apiHost.POST("/update_ext", jwt.AuthAdmin, host.UpdateHostExt)
	apiHost.POST("/update_add_token", jwt.AuthAdmin, host.UpdateHostAddToken)
	apiHost.POST("/get_add_token", jwt.AuthAdmin, host.GetHostAddToken)
	apiHost.POST("/delete", jwt.AuthAdmin, host.DeleteHost)
	apiHost.POST("/authorized_list", jwt.AuthLogin, host.GetHostAuthorizedList)
	apiHost.POST("/status", jwt.AuthLogin, host.GetHostStatus)
	apiHost.POST("/image_list", jwt.AuthLogin, host.GetHostImageList)
	apiHost.POST("/profile", jwt.AuthLogin, host.GetHostProfile)

	apiContainer := apiG.Group("/container")
	apiContainer.POST("/create", jwt.AuthLogin, container.ContainerCreate)

	apiOperation := apiG.Group("/operation")
	apiOperation.POST("/status", jwt.AuthLogin, operation.OperationStatus)

	callbackG := r.Group("/callback")
	callbackHost := callbackG.Group("/host")
	callbackHost.POST("/delete_callback", host.DeleteHostCallback)
	callbackContainer := callbackG.Group("/container")
	callbackContainer.POST("/create_callback", container.ContainerCreateCallback)

	return r
}
