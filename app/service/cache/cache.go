package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"gigame.xyz/app/vars"
	libredis "gigame.xyz/library/redis"
	"gorm.io/gorm"
	"strings"
	"time"
)

type DBCache struct {
	db          *gorm.DB
	cachePrefix string
	driver      *libredis.Redis
	ctx         context.Context
	debug       bool
	expire      time.Duration
}

func NewCache(db *gorm.DB) *DBCache {
	return &DBCache{
		cachePrefix: vars.RedisKeyPrefix,
		driver:      vars.Redis,
		db:          db,
		ctx:         context.Background(),
		expire:      time.Second * 7200,
		debug:       vars.YmlConfig.GetBool("Debug"),
	}
}

// QueryRows 条件查找 v 必需是指针类型参数
func (dc *DBCache) QueryRows(k string, v interface{}, f func(_db *gorm.DB, _v interface{}, _bys ...interface{}) error, bys ...interface{}) (err error) {
	//pv := reflect.ValueOf(v)
	//if pv.Kind() != reflect.Ptr {
	//
	//}
	if dc.debug {
		return f(dc.db, v, bys...)
	}
	var keys int64 = 0
	cacheKey := dc.autoBuildKey(k, bys...)
	if keys, err = dc.driver.Exists(dc.ctx, cacheKey).Result(); err != nil {
		return err
	} else if keys == 0 {
		if err = f(dc.db, v, bys...); err != nil {
			return err
		}
		_ = dc.setCache(cacheKey, v)
		return nil
	}
	result, err := dc.driver.Get(dc.ctx, cacheKey).Bytes()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(result, v); err != nil {
		return err
	}
	return nil
}

func (dc *DBCache) RemoveBy(key string, by interface{}) (err error) {
	_, err = dc.driver.Del(dc.ctx, dc.autoBuildKey(key, by)).Result()
	return err
}

func (dc *DBCache) setCache(key string, v interface{}) (err error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = dc.driver.Set(dc.ctx, key, string(marshal), dc.expire).Result()
	return err
}

func (dc *DBCache) autoBuildKey(key string, ks ...interface{}) string {
	s := []string{"%s:%s"}
	vs := []interface{}{dc.cachePrefix, key}
	for _, k := range ks {
		s = append(s, "%v")
		vs = append(vs, k)
	}
	return fmt.Sprintf(strings.Join(s, ":"), vs...)
}

// SetExpire 手动设置过期时间，单位 秒
func (dc *DBCache) SetExpire(sec int) {
	dc.expire = time.Second * time.Duration(sec)
}
