package etcd

import (
	"encoding/json"
	"time"

	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"github.com/astaxie/beego"
	"github.com/2young2simple/YiJob/model"
)

type Cluster struct {
	nodes   []*model.Node
	KeysAPI client.KeysAPI
}

func NewCluster(endpoints []string) (*Cluster, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	etcdClient, err := client.New(cfg)
	if err != nil {
		beego.Error("Error: cannot connec to etcd:", err)
		return nil, err
	}

	master := &Cluster{
		nodes:   []*model.Node{},
		KeysAPI: client.NewKeysAPI(etcdClient),
	}
	return master, nil
}

func (c *Cluster) Start() error {
	go c.WatchWorkers()
	go c.loggerNodes()
	fmt.Println("Master Start ...")
	return nil
}

func (c *Cluster) GetNodes() []*model.Node {
	return c.nodes
}

func (c *Cluster) addWorker(info *model.Node) {
	node := &model.Node{
		IsHealth:   true,
		IP:         info.IP,
		Port:       info.Port,
		Name:       info.Name,
	}

	n := c.getNode(node.Name)
	if n == nil{
		c.nodes = append(c.nodes,node)
	}else{
		n.IsHealth = true
		n.IP = info.IP
		n.Port = info.Port
	}
}

func (c *Cluster)getNode(name string) *model.Node{
	for i,node := range c.nodes{
		if node.Name == name{
			return c.nodes[i]
		}
	}
	return nil
}

func (c *Cluster)deleteNode(name string) {
	for i,node := range c.nodes{
		if node.Name == name{
			c.nodes = append(c.nodes[:i],c.nodes[i+1:]...)
			return
		}
	}
}

func (c *Cluster) WatchWorkers() {
	api := c.KeysAPI
	watcher := api.Watcher("nodes/", &client.WatcherOptions{
		Recursive: true,
	})
	for {
		res, err := watcher.Next(context.Background())
		if err != nil {
			beego.Error("Error watch workers:", err)
			break
		}
		if res.Action == "expire" {
			info,err := unmarshal(res.PrevNode)
			if err != nil{
				beego.Error("解析消息失败：",err.Error())
				continue
			}

			beego.Warning("Agent 失去消息 ：", info.Name)
			node := c.getNode(info.Name)
			if node != nil {
				node.IsHealth = false
			}
		} else if res.Action == "set" {
			info,err := unmarshal(res.Node)
			if err != nil{
				beego.Error("解析消息失败：",err.Error())
				continue
			}
			c.addWorker(info)
		} else if res.Action == "delete" {
			info,err := unmarshal(res.Node)
			if err != nil{
				beego.Error("解析消息失败：",err.Error())
				continue
			}
			beego.Warning("删除 Agent ：", info.Name)
			c.deleteNode(info.Name)

		}
	}
}

func (c *Cluster)loggerNodes(){
	for {
		nodes := []model.Node{}
		for _,node := range c.nodes{
			nodes = append(nodes,*node)
		}
		beego.Info("节点：",nodes)
		time.Sleep(10*time.Second)
	}

}

func unmarshal(node *client.Node) (*model.Node,error) {
	info := &model.Node{}
	err := json.Unmarshal([]byte(node.Value), info)
	if err != nil {
		return nil,err
	}
	return info,nil
}
