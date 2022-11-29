package global

import (
	"sync"

	"github.com/ciscolive/gin-admin/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/ciscolive/gin-admin-common/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	DBList      map[string]*gorm.DB
	Redis       *redis.Client
	Config      config.Server
	Resty       *resty.Client
	Viper       *viper.Viper
	Logger      *zap.Logger // Logger    *oplogging.Logger
	Timer       = timer.NewTimerTask()
	Concurrency = &singleflight.Group{}
	BlackCache  local_cache.Cache
	lock        sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("GetGlobalDBByDBName panic")
	}
	return db
}
