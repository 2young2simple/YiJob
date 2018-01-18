package main

import (
	"testing"
	"github.com/astaxie/beego/httplib"
	"github.com/2young2simple/YiJob/model"
	"fmt"
)

func Test_AddJob(t *testing.T){
	job := model.JobModel{Name:"print",Params:"",Status:1}
	req,err := httplib.Post(`http://127.0.0.1:6060/api/job`).JSONBody(job)
	if err != nil{
		t.Fatal(err)
	}
	result,err := req.Bytes()
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(string(result))
}
