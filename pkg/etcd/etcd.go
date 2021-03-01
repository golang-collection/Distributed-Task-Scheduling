package etcd

import (
	"Distributed-Task-Scheduling/pkg/setting"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-06 18:22
* @Description:
**/

func NewEtcdEngine(etcdSetting *setting.EtcdSettingS) (client *clientv3.Client, kv clientv3.KV, lease clientv3.Lease, watcher clientv3.Watcher, err error) {
	config := clientv3.Config{
		Endpoints:            []string{etcdSetting.Endpoint},
		DialTimeout:          time.Duration(etcdSetting.DialTimeout) * time.Millisecond,
		DialKeepAliveTime:    time.Duration(etcdSetting.DialKeepAliveTime) * time.Second,
		DialKeepAliveTimeout: time.Duration(etcdSetting.DialKeepAliveTimeout) * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		return
	}
	if client == nil{
		err = errors.New("init etcd error")
		return
	}
	if _, err = client.Dial(etcdSetting.Endpoint); err != nil{
		return
	}
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)
	return
}
