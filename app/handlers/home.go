package handlers

import (
	"game.mobgi.cc/app/response"
	"game.mobgi.cc/app/vars"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Home struct{}

func (h Home) Home(ctx *gin.Context) {
	games := make(map[string][]*vars.Game)
	for _, game := range vars.Games {
		if game.Tag == vars.HotGame {
			games[vars.HotGame] = append(games[vars.HotGame], game)
		} else if game.Tag == vars.NewGame {
			games[vars.NewGame] = append(games[vars.NewGame], game)
		} else {
			games[vars.NormalGame] = append(games[vars.NormalGame], game)
		}
	}
	response.HTML(ctx, "index", gin.H{"games": games})
}

func (h Home) GamePage(ctx *gin.Context) {
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

	response.HTML(ctx, "game", gin.H{"game": game, "games": showGames})
}

func (h Home) PlayTrans(ctx *gin.Context) {
	id := ctx.Query("gid")
	if id == "" {
		ctx.Redirect(http.StatusMovedPermanently, "/")
	}
	response.HTML(ctx, "full-screen-ad", gin.H{"gid": id})
}

func (h Home) GamePlay(ctx *gin.Context) {
	id := ctx.Query("gid")
	var game *vars.Game
	for _, g := range vars.Games {
		if g.GameId == id {
			game = g
		}
	}

	response.HTML(ctx, "play", gin.H{"game": game})
}
