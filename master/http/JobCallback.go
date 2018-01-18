package http

import (
	"github.com/gin-gonic/gin"
	"github.com/2young2simple/YiJob/model"
	"github.com/2young2simple/YiJob/master/db"
	"time"
)

func JobCallback(ctx *gin.Context){
	job := &model.JobModel{}
	if err := ctx.BindJSON(job);err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}
	job.EndTime = time.Now()
	result,err := db.UpdateJob(job)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	WriteJson(ctx,1,"success",result)
	return
}
