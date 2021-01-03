package global

import (
	"github.com/jinzhu/gorm"
	"myblog/pkg/logger"
	"myblog/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettings
	AppSetting *setting.AppSettings
	DataBaseSetting *setting.DatabaseSettings
	DBEngine *gorm.DB
	Logger *logger.Logger
)


