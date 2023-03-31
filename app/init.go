package app

import (
	"game.mobgi.cc/app/vars"
	"game.mobgi.cc/library/config"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func init() {
	checkConfigFiles()

	// 初始化 WEB 配置文件
	vars.YmlConfig = config.CreateYamlFactory()
	vars.YmlConfig.ConfigFileChangeListen()

	loadIniGames()
}

// 检查必要的配置文件
func checkConfigFiles() {
	if _, err := os.Stat(vars.BasePath + "/config/web.yaml"); err != nil {
		log.Fatal("请检查 WEB 配置文件是否存在：", err)
		return
	}
	if _, err := os.Stat(vars.BasePath + "/config/games.ini"); err != nil {
		log.Fatal("请检查 游戏 配置文件是否存在：", err)
		return
	}
}

func loadIniGames() {
	cfg, err := ini.Load(vars.BasePath + "/config/games.ini")
	if err != nil {
		log.Fatal("游戏配置文件加载失败", err)
		return
	}
	for _, section := range cfg.Sections() {
		if section.Name() == "" || section.Name() == "DEFAULT" {
			continue
		}
		gameMap := section.KeysHash()
		vars.Games = append(vars.Games, &vars.Game{
			Name:        section.Name(),
			GameId:      gameMap["GameId"],
			GameUrl:     gameMap["GameUrl"],
			Icon:        gameMap["Icon"],
			Description: gameMap["Description"],
			Tag:         gameMap["Tag"],
			ShowAd:      0,
		})
	}
}
