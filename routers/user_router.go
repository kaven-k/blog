package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"server/api"
	"server/middleware"
)

var store = cookie.NewStore([]byte("hygafewf154156asfdewaf"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionID", store))
	router.POST("email_login", userApi.EmailLoginView)
	router.POST("login", userApi.QQLoginView)
	router.POST("users_create", userApi.UserCreateView) // middleware.JwtAdmin(),
	router.GET("users", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("user_update_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	router.PUT("user_update_password", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	router.POST("logout", middleware.JwtAuth(), userApi.UserLogoutView)
	router.DELETE("users_delete", middleware.JwtAdmin(), userApi.UserRemoveView)
	router.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)

}
