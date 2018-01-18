package job

import (
	"github.com/2young2simple/YiJob/model"
	"errors"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

var (
	WorkerI Worker
	masterAddr string
)

const JobCHanSize  = 100

func init(){
	WorkerI = Worker{Jobs:make(map[string]Job),JobChan:make(chan model.JobModel,JobCHanSize)}
	masterAddr = beego.AppConfig.String("master_addr")
}

type Worker struct {
	Jobs map[string]Job
	JobChan chan model.JobModel
}

func (w *Worker) Run(){
	for {
		jobModel := <- WorkerI.JobChan
		job,ok := w.Jobs[jobModel.Name]
		if !ok{
			return
		}
		w.doJob(job,jobModel)
	}
}

func (w *Worker) doJob(job Job,jm model.JobModel) {
	defer job.Finish(jm)
	result,err := job.Do(jm)
	if err != nil{
		jobFail(jm,"err:"+err.Error()+" result :"+result)
		return
	}
	jobSuccess(jm,result)
}

func (w *Worker) AddJob(name string,job Job){
	if _,ok := w.Jobs[name];ok{
		panic("任务已存在:"+name)
	}
	w.Jobs[name] = job
}

func (w *Worker) IsExistJob(name string) bool{
	_,ok := w.Jobs[name]
	return ok
}

func (w *Worker)Push(m model.JobModel) error{
	if len(w.JobChan) >= JobCHanSize{
		return errors.New("添加任务失败，任务队列已满")
	}
	w.JobChan <- m
	return nil
}


func jobSuccess(job model.JobModel,result string){
	job.Status = model.Pub_Status_Success
	job.Result = result
	callback(job)
}

func jobFail(job model.JobModel,result string){
	job.Status = model.Pub_Status_Fail
	job.Result = result
	callback(job)

}

func callback(job model.JobModel){
	url := fmt.Sprintf(`http://%s/api/job/callback`,masterAddr)
	result := model.Result{}
	req,err := httplib.Post(url).JSONBody(job)
	if err != nil{
		beego.Error(err.Error())
		return
	}
	err = req.ToJSON(&result)
	if err != nil{
		beego.Error(err.Error())
		return
	}
	beego.Info("job:",job.Name," callback:",result)
}


