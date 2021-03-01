package initConf

import (
	"Distributed-Task-Scheduling/global"
	"Distributed-Task-Scheduling/pkg/cache"
	"Distributed-Task-Scheduling/pkg/db"
	"Distributed-Task-Scheduling/pkg/etcd"
	"Distributed-Task-Scheduling/pkg/logger"
	"Distributed-Task-Scheduling/pkg/mongoDB"
	"Distributed-Task-Scheduling/pkg/setting"
	"log"
	"strings"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

/**
* @Author: super
* @Date: 2021-01-05 14:25
* @Description:
**/
func Init(config string) {
	//初始化配置
	err := setupSetting(config)
	if err != nil {
		log.Printf("init setupSetting err: %v\n", err)
	} else {
		log.Printf("初始化配置信息成功")
	}
	//初始化日志
	err = setupLogger()
	if err != nil {
		log.Printf("init setupLogger err: %v\n", err)
	} else {
		log.Printf("初始化logger成功")
	}
	//初始化数据库
	err = setupDBEngine()
	if err != nil {
		log.Printf("init setupDBEngine err: %v\n", err)
	} else {
		log.Printf("初始化数据库成功")
	}
	//初始化redis
	err = setupCacheEngine()
	if err != nil {
		log.Printf("init setupCacheEngine err: %v\n", err)
	} else {
		log.Printf("初始化cache成功")
	}
	//初始化mongoDB
	err = setupMongoDBEngine()
	if err != nil {
		log.Printf("init setupMongoDBEngine err: %v\n", err)
	} else {
		log.Printf("初始化mongoDb成功")
	}
	//初始化etcd
	err = setupEtcdEngine()
	if err != nil {
		log.Printf("init setupEtcdEngine err: %v\n", err)
	} else {
		log.Printf("初始化etcd成功")
	}
}

func setupSetting(config string) error {
	newSetting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("MongoDB", &global.MongoDBSetting)
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Etcd", &global.EtcdSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = db.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupCacheEngine() error {
	var err error
	global.RedisEngine, err = cache.NewRedisEngine(global.CacheSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupMongoDBEngine() error {
	var err error
	global.MongoDBEngine, err = mongoDB.NewMongoDBEngine(global.MongoDBSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupEtcdEngine() error {
	var err error
	global.EtcdEngine, global.EtcdKV, global.EtcdLease, global.EtcdWatcher, err = etcd.NewEtcdEngine(global.EtcdSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	log.Println("log file name ", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
