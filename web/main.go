package main

import (
	_ "game.mobgi.cc/app"
	"game.mobgi.cc/router"
	"log"
)

func main() {
	if err := router.Router(); err != nil {
		log.Fatal("app start failed:", err)
	}
}
