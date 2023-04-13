package app

import (
	"gigame.xyz/app/vars"
	"gigame.xyz/library/config"
	libmysql "gigame.xyz/library/gorm/mysql"
	libredis "gigame.xyz/library/redis"
	"log"
	"os"
	"time"
)

func init() {
	checkConfigFiles()

	// 初始化 WEB 配置文件
	vars.YmlConfig = config.CreateYamlFactory()
	vars.YmlConfig.ConfigFileChangeListen()

	initDb()
	initRedis()
}

// 检查必要的配置文件
func checkConfigFiles() {
	if _, err := os.Stat(vars.BasePath + "/config/web.yaml"); err != nil {
		log.Fatal("请检查 WEB 配置文件是否存在：", err)
		return
	}
}

func initDb() {
	addr := vars.YmlConfig.GetString("Mysql.Host")
	db := vars.YmlConfig.GetString("Mysql.Database")
	username := vars.YmlConfig.GetString("Mysql.Username")
	pwd := vars.YmlConfig.GetString("Mysql.Password")
	charset := vars.YmlConfig.GetString("Mysql.Charset")
	port := vars.YmlConfig.GetInt("Mysql.Port")
	life := vars.YmlConfig.GetInt("Mysql.SetConnMaxLifetime")
	if mysql, err := libmysql.NewMysql(addr, username, pwd, db, charset, port, time.Second*time.Duration(life)); err != nil {
		log.Fatal("Redis 连接失败", err.Error())
		return
	} else {
		vars.Mysql = mysql
	}
}

func initRedis() {
	addr := vars.YmlConfig.GetString("Redis.Host")
	pwd := vars.YmlConfig.GetString("Redis.Password")
	db := vars.YmlConfig.GetInt("Redis.Db")
	if redis, err := libredis.NewRedis(addr, pwd, vars.RedisKeyPrefix, db); err != nil {
		log.Fatal("Redis 连接失败", err.Error())
		return
	} else {
		vars.Redis = redis
	}
}
