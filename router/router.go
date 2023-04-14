package router

import (
	"gigame.xyz/app/handlers/box1"
	"gigame.xyz/app/middleware"
	"gigame.xyz/app/validators"
	"net/http"

	"gigame.xyz/app/vars"
	"github.com/gin-gonic/gin"
)

func Router() error {
	r := gin.Default()

	//r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("assets", http.Dir("./assets"))
	//address := vars.YmlConfig.GetString("Redis.Host")
	//pass := vars.YmlConfig.GetString("Redis.Password")
	//store, _ := redis.NewStore(10, "tcp", address, pass, []byte(""))
	//r.Use(sessions.Sessions("GiGameSessions", store))

	if vars.YmlConfig.GetBool("HttpServer.AllowCrossDomain") {
		r.Use(corsNext())
	}

	g := r.Group("/api/box01", middleware.CheckAuth())
	{
		g.GET("/ad-config", (box1.Api{}).AdConfig)
		g.GET("/games", (box1.Api{}).Games)
		g.GET("/game/:gid", (box1.Api{}).GameInfo)
		g.GET("/search", (box1.Api{}).Search)
		g.GET("/topic", (validators.V{}).Topic)
	}

	return r.Run(vars.YmlConfig.GetString("HttpServer.Port"))
}

// 允许跨域
func corsNext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Authorization,User-Agent, Keep-Alive, Content-Type, X-Requested-With,X-CSRF-Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusAccepted)
		}
		c.Next()
	}
}
