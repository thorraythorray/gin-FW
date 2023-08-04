package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/thorraythorray/go-proj/config"
	"github.com/thorraythorray/go-proj/ginx/schema"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mode    string
	Confile string
	Config  *config.ConfigMap
	Redis   *redis.Client
	DB      *gorm.DB
	Logger  *zap.SugaredLogger
	User    *schema.AuthorizedUserInfo
)
