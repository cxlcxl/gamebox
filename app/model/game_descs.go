package model

import (
	"gigame.xyz/app/service/cache"
	"gorm.io/gorm"
)

type GameDesc struct {
	db

	GameId string `json:"game_id" gorm:"primaryKey"`
	Desc   string `json:"desc" gorm:"column:_desc"`
}

func (m *GameDesc) TableName() string {
	return "game_descs"
}

func DbGameDesc(_db *gorm.DB) *GameDesc {
	return &GameDesc{db: db{DB: _db}}
}

func (m *GameDesc) FindDescById(gid string) (desc *GameDesc, err error) {
	err = cache.NewCache(m.DB).QueryRows(cacheDbKey, &desc, func(_db *gorm.DB, v interface{}, i ...interface{}) error {
		return _db.Table(i[0].(string)).Where("game_id = ?", i[1]).First(v).Error
	}, m.TableName(), gid)
	return
}
