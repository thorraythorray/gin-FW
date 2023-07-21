package global

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/thorraythorray/go-proj/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Mode    string
	Confile string
	Config  *config.ConfigMap
	Redis   *redis.Client
	DB      *gorm.DB
	DBPool  *sql.DB
	Logger  *zap.SugaredLogger
)
