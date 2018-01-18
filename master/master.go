package master

import (
	"github.com/astaxie/beego"
	"github.com/2young2simple/YiJob/master/discover"
	"github.com/2young2simple/YiJob/master/dispatch"
	"github.com/2young2simple/YiJob/master/http"
)

func Start(){
	etcdAddr := beego.AppConfig.String("etcd")

	if err := discover.InitDiscover(etcdAddr);err != nil{
		panic(err)
	}
	go http.ServerStart()
	go dispatch.Run()
}
