package global

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
* @Author: super
* @Date: 2020-09-18 08:51
* @Description: 全局配置DB
**/

var (
	DBEngine      *gorm.DB
	RedisEngine   *redis.Pool
	MongoDBEngine *mongo.Client
	EtcdEngine    *clientv3.Client
	EtcdKV        clientv3.KV
	EtcdLease     clientv3.Lease
	EtcdWatcher   clientv3.Watcher
)
