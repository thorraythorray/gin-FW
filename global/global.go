package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/thorraythorray/go-proj/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mode       string
	ConfigPath string
	ConfigData *config.ConfigMap
	Redis      *redis.Client
	DB         *gorm.DB
	Logger     *zap.SugaredLogger
)
