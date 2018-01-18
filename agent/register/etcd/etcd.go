package etcd

import (
	"encoding/json"
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"github.com/2young2simple/YiJob/model"
	"fmt"
)

func Register(node model.Node,endpoints []string){
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		panic(err)
	}

	go heartBeat(c,node)

	fmt.Println("Agent Start ...")
}

func heartBeat(c client.Client,node model.Node) {
	api := client.NewKeysAPI(c)
	for {

		key := "nodes/" + node.Name
		value, _ := json.Marshal(node)

		_, err := api.Set(context.Background(), key, string(value), &client.SetOptions{
			TTL: time.Second * 15,
		})
		if err != nil {
			log.Println("Error update workerInfo:", err)
		}
		time.Sleep(time.Second * 5)
	}
}
