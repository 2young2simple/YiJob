package agent

import (
	"github.com/astaxie/beego"
	"github.com/2young2simple/YiJob/model"
	"github.com/2young2simple/YiJob/agent/register/etcd"
	"github.com/2young2simple/YiJob/agent/job"
	"github.com/2young2simple/YiJob/agent/http"
)

func AddJob(name string,j job.Job){
	job.WorkerI.AddJob(name,j)
}

func Run(){
	etcdAddr := beego.AppConfig.String("etcd")
	name := beego.AppConfig.String("name")
	ip := beego.AppConfig.String("ip")
	port := beego.AppConfig.String("port")

	etcd.Register(model.Node{Name:name,IP:ip,Port:port},[]string{etcdAddr})

	go http.ServerStart()
	go job.WorkerI.Run()
}
