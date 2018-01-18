package db

import (
	"github.com/2young2simple/YiJob/model"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func InsertJob(job *model.JobModel) (*model.JobModel,error){
	o := orm.NewOrm()
	if _,err := o.Insert(job);err != nil{
		return nil,err
	}
	return job,nil
}

func UpdateJob(job *model.JobModel) (*model.JobModel,error){
	o := orm.NewOrm()
	if _,err := o.Update(job,"begin_time","end_time","status","result");err != nil{
		return nil,err
	}
	return job,nil
}

func ListJobs() ([]model.JobModel,error){
	o := orm.NewOrm()
	sql := `select * from job_model`
	jobs := []model.JobModel{}
	_,err := o.Raw(sql).QueryRows(&jobs)
	if err != nil{
		return nil,err
	}
	return jobs,nil
}

func GetJobs(id int) (*model.JobModel,error){
	o := orm.NewOrm()
	sql := fmt.Sprintf("select * from job_model where id = %d", id)
	job := &model.JobModel{}
	err := o.Raw(sql).QueryRow(job)
	if err != nil {
		return nil,err
	}
	return job,err
}

func DeteleJob(id int) error{
	o := orm.NewOrm()
	if _,err := o.Delete(model.JobModel{Id:id},"id");err != nil{
		return err
	}
	return nil
}

func GetPreJobs() ([]model.JobModel,error){
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil{
		o.Rollback()
		return nil,err
	}

	pubs := []model.JobModel{}
	_,err = o.Raw("select * from job_model where status = ?",model.Pub_Status_Prepub).QueryRows(&pubs)
	if err != nil{
		o.Rollback()
		return nil,err
	}
	for i := range pubs{
		pub := pubs[i]
		pub.Status = model.Pub_Status_Pubing
		_,err := o.Update(&pub,"status")
		if err != nil{
			o.Rollback()
			return nil,err
		}
	}

	err = o.Commit()
	if err != nil{
		o.Rollback()
		return nil,err
	}

	return pubs,nil
}