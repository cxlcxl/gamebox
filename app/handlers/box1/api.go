package box1

import (
	"gigame.xyz/app/response"
	"gigame.xyz/app/vars"
	"github.com/gin-gonic/gin"
	"strings"
)

type Api struct{}

func (h Api) Games(ctx *gin.Context) {
	games := make(map[string][]*vars.Game)
	for _, game := range vars.Games {
		if game.Tag == vars.HotGame {
			games[vars.HotGame] = append(games[vars.HotGame], game)
		}
		if game.Tag == vars.NewGame {
			games[vars.NewGame] = append(games[vars.NewGame], game)
		}
		games["All"] = append(games["All"], game)
	}
	response.Success(ctx, games)
}

func (h Api) GameInfo(ctx *gin.Context) {
	id := ctx.Query("gid")
	var game *vars.Game
	var showGames []*vars.Game
	for _, g := range vars.Games {
		if g.GameId == id {
			game = g
		}
		if g.Tag == vars.HotGame && g.GameId != id {
			showGames = append(showGames, g)
		}
	}

	response.Success(ctx, gin.H{"game": game, "games": showGames})
}

func (h Api) GamePlay(ctx *gin.Context) {
	gid := ctx.Param("gid")
	var game *vars.Game
	for _, g := range vars.Games {
		if g.GameId == gid {
			game = g
			break
		}
	}

	response.Success(ctx, game)
}

func (h Api) Search(ctx *gin.Context) {
	k := ctx.Query("k")
	var games []*vars.Game
	for _, g := range vars.Games {
		if strings.Contains(g.Name, k) {
			games = append(games, g)
		}
	}

	response.Success(ctx, games)
}
