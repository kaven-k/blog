package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	// qq登录部分测试，暂时先写在这里
	//router.GET("login", user_api.UserApi{}.QQLoginView)

	// 路由分组
	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层
	// 系统配置API
	routerGroupApp.SettingsRouter()
	// 图片上传API
	routerGroupApp.ImagesRouter()
	// 广告API
	routerGroupApp.AdvertisementRouter()
	// 菜单API
	routerGroupApp.MenuRouter()
	// 用户API
	routerGroupApp.UserRouter()
	// 标签API
	routerGroupApp.TagRouter()
	// 消息API
	routerGroupApp.MessageRouter()
	// 文章API
	routerGroupApp.ArticleRouter()
	// 评论点赞API
	routerGroupApp.DiggRouter()
	return router
}
