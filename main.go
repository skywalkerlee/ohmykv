package main

import (
	"net"

	"os"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/cluster"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/skywalkerlee/ohmykv/msg"
	"github.com/skywalkerlee/ohmykv/rpc"
	"google.golang.org/grpc"
)

func main() {
	cluster := cluster.NewCluster()
	go cluster.Sync()
	lis, err := net.Listen("tcp", config.Ohmkvcfg.Ohmkv.Addr+":"+config.Ohmkvcfg.Ohmkv.Port)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	s := grpc.NewServer()
	msg.RegisterKvServer(s, &rpc.Service{Cluster: cluster})
	s.Serve(lis)
}
