package http

import (
	"github.com/gin-gonic/gin"
	"github.com/2young2simple/YiJob/model"
	"github.com/2young2simple/YiJob/agent/job"
)

func DoJob(ctx *gin.Context){
	jobModel := model.JobModel{}
	err := ctx.BindJSON(&jobModel)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	if !job.WorkerI.IsExistJob(jobModel.Name){
		WriteJson(ctx,-1,"任务不存在："+jobModel.Name,nil)
		return
	}

	err = job.WorkerI.Push(jobModel)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	WriteJson(ctx,1,"任务已加入队列",nil)
	return
}
