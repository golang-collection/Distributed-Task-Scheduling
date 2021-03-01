package main

import (
	"Distributed-Task-Scheduling/initConf"
	"Distributed-Task-Scheduling/internal/crontab/worker"
	"context"
	"log"
	"time"
)

/**
* @Author: super
* @Date: 2021-02-08 23:04
* @Description:
**/

func main() {
	initConf.Init("/Users/super/develop/Distributed-Task-Scheduling/configs")
	if err := worker.NewScheduler(); err != nil {
		log.Printf("init NewScheduler err: %v\n", err)
		return
	}
	if err := worker.NewExecutor(); err != nil {
		log.Printf("init NewExecutor err: %v\n", err)
		return
	}
	if err := worker.NewLogSink(); err != nil {
		log.Printf("init NewLogSink err: %v\n", err)
		return
	}
	if err := worker.WatchJobs(context.Background()); err != nil {
		log.Printf("init WatchJobs err: %v\n", err)
		return
	}
	worker.WatchKiller(context.Background())
	go worker.KeepOnline()

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

}
