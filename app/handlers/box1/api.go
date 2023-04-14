package box1

import (
	"gigame.xyz/app/model"
	"gigame.xyz/app/response"
	"gigame.xyz/app/validators/v_data_01"
	"gigame.xyz/app/vars"
	"github.com/gin-gonic/gin"
	"strings"
)

type Api struct{}

func (h Api) AdConfig(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"is_show": 1,
	})
}

func (h Api) Games(ctx *gin.Context) {
	games, err := model.DbGame(vars.Mysql).FetchGamesByBox(vars.GameBox01)
	if err != nil {
		response.Fail(ctx)
		return
	}
	_games := make(map[string][]*model.ListGame)
	for _, game := range games {
		if strings.Contains(game.Topics, vars.HotGame) {
			_games[vars.HotGame] = append(_games[vars.HotGame], game)
		}
		if strings.Contains(game.Topics, vars.NewGame) {
			_games[vars.NewGame] = append(_games[vars.NewGame], game)
		}
		_games["All"] = append(_games["All"], game)
	}
	response.Success(ctx, _games)
}

func (h Api) GameInfo(ctx *gin.Context) {
	gid := ctx.Param("gid")
	game, err := model.DbGame(vars.Mysql).FindGameById(vars.GameBox01, gid)
	if err != nil {
		response.Fail(ctx)
		return
	}
	if desc, err := model.DbGameDesc(vars.Mysql).FindDescById(gid); err == nil {
		game.Desc = desc.Desc
	}
	response.Success(ctx, game)
}

func (h Api) Search(ctx *gin.Context) {
	k := ctx.Query("k")
	searchGames, err := model.DbGame(vars.Mysql).SearchGames(vars.GameBox01, strings.TrimSpace(k))
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.Success(ctx, searchGames)
}

func (h Api) Topic(ctx *gin.Context, p interface{}) {
	params := p.(*v_data_01.VDataTopic)
	games, err := model.DbGame(vars.Mysql).FindGamesByTag(vars.GameBox01, params.Topic)
	if err != nil {
		response.Fail(ctx)
		return
	}

	response.Success(ctx, games)
}
