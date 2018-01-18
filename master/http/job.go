package http

import (
	"github.com/gin-gonic/gin"
	"github.com/2young2simple/YiJob/model"
	"github.com/2young2simple/YiJob/master/db"
	"strconv"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/2young2simple/YiJob/master/dispatch"
)

func AddJob(ctx *gin.Context){

	job := &model.JobModel{}
	if err := ctx.BindJSON(job);err != nil{
		fmt.Println("err:",err)
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}
	job.Status = model.Pub_Status_Pubing
	result,err := db.InsertJob(job)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}
	beego.Info("创建任务成功：",result)
	dispatch.Push(*result)

	WriteJson(ctx,1,"success",result)
	return
}

func GetJob(ctx *gin.Context){
	idStr := ctx.Query("id")
	id,err := strconv.Atoi(idStr)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	result,err := db.GetJobs(id)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	WriteJson(ctx,1,"success",result)
	return
}

func ListJob(ctx *gin.Context){
	result,err := db.ListJobs()
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	WriteJson(ctx,1,"success",result)
	return
}

func DeleteJob(ctx *gin.Context){
	type Id struct {
		Id int `json:"id"`
	}
	id := &Id{}
	if err := ctx.BindJSON(id);err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	err := db.DeteleJob(id.Id)
	if err != nil{
		WriteJson(ctx,-1,err.Error(),nil)
		return
	}

	WriteJson(ctx,1,"success",nil)
	return
}