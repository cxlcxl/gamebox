package model

import (
	"gigame.xyz/app/service/cache"
	"gorm.io/gorm"
)

type Game struct {
	db

	GameId string `json:"game_id" gorm:"uniqueIndex"`
	Name   string `json:"name" gorm:"column:_name"`
	Icon   string `json:"icon" gorm:"column:_icon"`
	Url    string `json:"url" gorm:"column:_url"`
	Topics string `json:"topics" gorm:"column:_topics"`
	Star   int    `json:"star" gorm:"column:_star"`
	Box    string `json:"box" gorm:"column:_box"`
}

var (
	listColumns = []string{"game_id", "_name", "_icon", "_star", "_topics"}
	infoColumns = []string{"game_id", "_name", "_icon", "_star", "_url"}
)

type ListGame struct {
	GameId string `json:"game_id"`
	Name   string `json:"name" gorm:"column:_name"`
	Icon   string `json:"icon" gorm:"column:_icon"`
	Topics string `json:"topics" gorm:"column:_topics"`
	Star   int    `json:"star" gorm:"column:_star"`
}

type InfoGame struct {
	GameId string `json:"game_id"`
	Name   string `json:"name" gorm:"column:_name"`
	Icon   string `json:"icon" gorm:"column:_icon"`
	Url    string `json:"url" gorm:"column:_url"`
	Star   int    `json:"star" gorm:"column:_star"`
	Desc   string `json:"desc" gorm:"-"`
}

func (m *Game) TableName() string {
	return "games"
}

func DbGame(_db *gorm.DB) *Game {
	return &Game{db: db{DB: _db}}
}

func (m *Game) FetchGamesByBox(box string) (games []*ListGame, err error) {
	err = cache.NewCache(m.DB).QueryRows(cacheDbKey, &games, func(_db *gorm.DB, v interface{}, i ...interface{}) error {
		return _db.Table(i[0].(string)).Select(listColumns).Where("_box = ?", i[1]).Find(v).Error
	}, m.TableName(), box)
	return
}

func (m *Game) FindGamesByTag(box, tag string) (games []*ListGame, err error) {
	err = cache.NewCache(m.DB).QueryRows(cacheDbKey, &games, func(_db *gorm.DB, v interface{}, i ...interface{}) error {
		return _db.Table(i[0].(string)).Select(listColumns).Where("_box = ? and find_in_set(?, _topics)", i[1], i[2]).Find(v).Error
	}, m.TableName(), box, tag)
	return
}

func (m *Game) FindGameById(box, gid string) (games *InfoGame, err error) {
	err = cache.NewCache(m.DB).QueryRows(cacheDbKey, &games, func(_db *gorm.DB, v interface{}, i ...interface{}) error {
		return _db.Table(i[0].(string)).Select(infoColumns).Where("_box = ? and game_id = ?", i[1], i[2]).First(v).Error
	}, m.TableName(), box, gid)
	return
}

func (m *Game) SearchGames(box, k string) (games []*ListGame, err error) {
	err = m.Table(m.TableName()).Select(listColumns).Where("_box = ? and _name like ?", box, "%"+k+"%").Find(&games).Error
	return
}
