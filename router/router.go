package router

import (
	"gigame.xyz/app/handlers/box1"
	"gigame.xyz/app/response"
	"net/http"

	"gigame.xyz/app/vars"
	"github.com/gin-gonic/gin"
)

func Router() error {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("assets", http.Dir("./assets"))

	//address := vars.YmlConfig.GetString("Redis.Host")
	//pass := vars.YmlConfig.GetString("Redis.Password")
	//store, _ := redis.NewStore(10, "tcp", address, pass, []byte(""))
	//r.Use(sessions.Sessions("GiGameSessions", store))

	//if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
	//	r.Use(corsNext())
	//}

	r.GET("/", (box1.Home{}).Home)
	r.GET("/game", (box1.Home{}).GamePage)
	r.GET("/play-trans", (box1.Home{}).PlayTrans)
	r.GET("/play", (box1.Home{}).GamePlay)
	r.GET("/search", (box1.Home{}).Search)

	r.GET("/about", func(ctx *gin.Context) {
		response.HTML(ctx, "about", nil)
	})
	r.GET("/privacy", func(ctx *gin.Context) {
		response.HTML(ctx, "privacy", nil)
	})
	r.GET("/cookies", func(ctx *gin.Context) {
		response.HTML(ctx, "cookie", nil)
	})
	r.GET("/contact", func(ctx *gin.Context) {
		response.HTML(ctx, "contact", nil)
	})
	r.GET("/copyright", func(ctx *gin.Context) {
		response.HTML(ctx, "copyright", nil)
	})

	r.GET("/ads.txt", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, vars.YmlConfig.GetString("AdsTxt"))
	})

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
