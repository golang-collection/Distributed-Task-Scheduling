package global

import (
	"Distributed-Task-Scheduling/pkg/logger"
	"Distributed-Task-Scheduling/pkg/setting"
)

/**
* @Author: super
* @Date: 2020-09-18 08:32
* @Description: 全局配置包括：服务，数据库，Email，JWT和日志
**/

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	EmailSetting    *setting.EmailSettingS
	Logger          *logger.Logger
	MongoDBSetting  *setting.MongoDBSettingS
	EtcdSetting     *setting.EtcdSettingS
)
