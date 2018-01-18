package dispatch

import (
	"github.com/2young2simple/YiJob/model"
	"errors"
)

var (
	Queue chan model.JobModel
)

func init(){
	Queue = make(chan model.JobModel,1000)
}

func Push(job model.JobModel) error{
	if len(Queue) >= 1000{
		return errors.New("分发队列已满")
	}
	Queue <- job
	return nil
}

func Pop() model.JobModel{
	return <- Queue
}
