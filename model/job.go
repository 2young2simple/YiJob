package model

import (
	"time"
)

const (
	Pub_Status_Nopub = 0
	Pub_Status_Prepub = 1
	Pub_Status_Pubing  = 2
	Pub_Status_Success  = 3

	Pub_Status_Fail  = -1
)

type JobModel struct {
	Id int `orm:"column(id);pk;auto" json:"id"`
	Name string `orm:"column(name)" json:"name"`
	Params string `orm:"column(params)" json:"params"`

	BeginTime time.Time `orm:"column(begin_time);null" json:"begin_time"`
	EndTime time.Time `orm:"column(end_time);null" json:"end_time"`
	Status int 	 `orm:"column(status)" json:"status"`   // 0 未执行  -1 发布失败  1 发布成功 -2 发送中

	Result string `orm:"column(result)" json:"result"`
}