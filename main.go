package main

import (
	"context"
	"net"

	"os"

	"flag"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/cluster"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/skywalkerlee/ohmykv/msg"
	"github.com/skywalkerlee/ohmykv/rpc"
	"google.golang.org/grpc"
)

func main() {
	leaderaddr := flag.String("leaderaddr", "", "host:port of leader to join")
	flag.Parse()
	if *leaderaddr != "" {
		conn, err := grpc.Dial(*leaderaddr, grpc.WithInsecure())
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(config.Ohmkvcfg.Raft.Addr + ":" + config.Ohmkvcfg.Ohmkv.Port)
		_, err = msg.NewKvClient(conn).Op(context.TODO(), &msg.Req{Op: 4, Value: []byte(config.Ohmkvcfg.Raft.Addr + ":" + config.Ohmkvcfg.Raft.Port)})
		if err != nil {
			logs.Error(err)
			return
		}
	}
	cluster := cluster.NewCluster(false)
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
