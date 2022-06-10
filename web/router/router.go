package router

import (
	"github.com/gin-gonic/gin"
	"ihome/web/conf"
	"ihome/web/controller"
	"ihome/web/model"
	"ihome/web/utils"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	//使用session
	// 初始化容器.
	//映射静态资源
	//fmt.Println(os.Getwd())
	//Chdir()
	//fmt.Println("111:")
	//fmt.Println(os.Getwd())
	//windows
	//router.Static("/home", "./web/view")

	//linux
	//router.Static("/home/login", "view/login.html")
	router.Static("/home", "view")

	v1 := router.Group("/api/v1.0")
	{

		v1.POST("/sessions", controller.PostLogin) //登录
		v1.GET("/imagecode/:uuid", controller.GetImageCd)
		v1.GET("/smscode/:phone", controller.GetSMSCd)
		v1.POST("/users", controller.Register)
		v1.GET("/session", controller.GetSession) //获取用户信息
		v1.GET("/areas", controller.GetAreas)     //获取房子地域信息

	}
	//router.Use(SessionAuthorize(router))
	//以下方法均使用session鉴权

	return router
}

func SessionAuthorize(r *gin.Engine) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		sessionId, err := ctx.Cookie(conf.LoginCookieName) // 获得cookie seesionId
		if err != nil || sessionId == "" {
			//cookie未存在或过期直接返回
			ctx.Abort()
			ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
			ctx.Request.URL.Path = "/home/login.html"
			r.HandleContext(ctx)
			return
		}
		//ts := c.Query("ts") // 时间戳
		//token := c.Query("token") // 访问令牌
		//session := sessions.Default(ctx)
		//user := session.Get(conf.SessionLoginIndex + "_" + sessionId)
		var user *model.User
		if user != nil {
			// 验证通过，会继续访问下一个中间件
			ctx.Next()
			return
		}
		// 验证不通过，不再调用后续的函数处理
		ctx.Abort()
		ctx.JSON(http.StatusOK, utils.Response(utils.RECODE_SESSIONERR, nil))
		ctx.Request.URL.Path = "/home/login.html"
		r.HandleContext(ctx)
		return
	}
}
