package routers

import (
	"Distributed-Task-Scheduling/global"
	"Distributed-Task-Scheduling/internal/middleware"
	"Distributed-Task-Scheduling/internal/routers/job"
	"Distributed-Task-Scheduling/internal/routers/sd"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
* @Author: super
* @Date: 2021-02-28 13:44
* @Description:
**/

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())

	r.StaticFile("/", "/Users/super/develop/Distributed-Task-Scheduling/internal/routers/web/index.html")
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	jobGroup := r.Group("/job")
	{
		jobGroup.POST("/save", job.SaveJob)
		jobGroup.POST("/delete", job.DeleteJob)
		jobGroup.GET("/list", job.ListJobs)
		jobGroup.POST("/kill", job.KillJob)
		jobGroup.GET("/log", job.JobLog)
	}
	r.GET("/worker/list", job.WorkerList)

	return r
}
