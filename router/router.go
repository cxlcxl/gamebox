package router

import (
	"game.mobgi.cc/app/handlers"
	"game.mobgi.cc/app/middleware"
	"net/http"

	"game.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
)

func Router() error {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("assets", http.Dir("./assets"))

	//address := vars.YmlConfig.GetString("Redis.Host")
	//pass := vars.YmlConfig.GetString("Redis.Password")
	//store, _ := redis.NewStore(10, "tcp", address, pass, []byte(""))
	//r.Use(sessions.Sessions("MobgiGameSessions", store))

	//if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
	//	r.Use(corsNext())
	//}

	r.GET("/", (handlers.Home{}).Home)
	r.GET("/game", (handlers.Home{}).GamePage)
	r.GET("/play-trans", (handlers.Home{}).PlayTrans)
	r.GET("/play", (handlers.Home{}).GamePlay)

	r.GET("/ads.txt", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, vars.YmlConfig.GetString("AdsTxt"))
	})

	r.Use(middleware.CheckUserLogin())
	{

	}
	//r.GET("/docs", func(ctx *gin.Context) {
	//ctx.HTML(http.StatusOK, "docs.tmpl", gin.H{})
	//})

	return r.Run(vars.YmlConfig.GetString("HttpServer.Port"))
}

// 允许跨域
func corsNext() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}