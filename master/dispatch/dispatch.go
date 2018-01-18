package dispatch

import (
	"github.com/2young2simple/YiJob/master/db"
	"github.com/2young2simple/YiJob/model"
	"time"
	"github.com/2young2simple/YiJob/master/discover"
	"github.com/astaxie/beego/httplib"
	"fmt"
	"github.com/astaxie/beego"
	"errors"
)

var dispatchCount = 1

func Run(){
	t := time.NewTicker(time.Second*10).C
	go dispatch()
	for range t{
		jobs,err := db.GetPreJobs()
		if err != nil{
			continue
		}
		for _,job := range jobs{
			Push(job)
		}
	}
}

func dispatch(){

	for {
		job := Pop()

		nodes := discover.GetNodes()

		if dispatchCount > 10000{
			dispatchCount = 1
		}

		if len(nodes) != 0{
			i := dispatchCount % len(nodes)
			if nodes[i].IsHealth{
				if err := nodeDo(nodes[i],job);err != nil{
					beego.Error("分发Agent任务 请求失败:",err)
					resetJob(job)
				}
				beego.Info("分发任务：",job," 节点：",nodes[i])

				beginJob(job)
				dispatchCount++
			}else{
				resetJob(job)
			}
		}
	}
}

func nodeDo(node *model.Node,job model.JobModel) error{
	url := fmt.Sprintf("http://%s:%s/api/job",node.IP,node.Port)
	result := &model.Result{}
	req,err := httplib.Post(url).JSONBody(job)
	if err != nil{
		return err
	}
	err = req.ToJSON(result)
	if err != nil{
		return err
	}

	if result.Code != 1{
		return errors.New(result.Message)
	}
	return nil
}

func resetJob(job model.JobModel){
	job.Status = model.Pub_Status_Prepub
	db.UpdateJob(&job)
}

func beginJob(job model.JobModel){
	job.BeginTime = time.Now()
	db.UpdateJob(&job)
}
