package global

import (
	config "GoToDoList/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config     *config.Config
	Logger     *zap.Logger
	Mysql      *gorm.DB
	Redis      *redis.Client
	GinContext *gin.Context
)
