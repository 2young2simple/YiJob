package main

import (
	"github.com/2young2simple/YiJob/agent"
	"github.com/astaxie/beego"
	"time"
	"github.com/2young2simple/YiJob/model"
	"github.com/kataras/go-errors"
)

type PrintJob struct {}

func (p *PrintJob) Do(jm model.JobModel) (string,error){
	beego.Info("开始执行输出任务...")
	time.Sleep(time.Second*2)
	beego.Info("输出任务执行完毕...")
	return "fail",errors.New("发生未知错误")
}

func (p *PrintJob) Finish(jm model.JobModel) {
	beego.Info("执行任务Finish回调")
}

func main() {
	agent.AddJob("print",&PrintJob{})
	agent.Run()
	select {}
}
