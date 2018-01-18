package job

import "github.com/2young2simple/YiJob/model"

type Job interface{
	Do(model.JobModel) (string,error)
	Finish(model.JobModel)
}