package discover

import (
	"github.com/2young2simple/YiJob/model"
	"github.com/2young2simple/YiJob/master/discover/etcd"
)

type Discover interface {
	GetNodes() []*model.Node
	Start() error
}

var discoverI Discover

func InitDiscover(etcdAddr string) error {
	var err error
	discoverI, err = etcd.NewCluster([]string{etcdAddr})
	if err != nil {
		return err
	}
	discoverI.Start()
	return nil
}

func GetNodes() []*model.Node {
	if discoverI != nil {
		return discoverI.GetNodes()
	}
	return nil
}
