package main

import (
	_ "gigame.xyz/app"
	"gigame.xyz/router"
	"log"
)

func main() {
	if err := router.Router(); err != nil {
		log.Fatal("app start failed:", err)
	}
}
