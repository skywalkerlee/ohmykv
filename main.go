package main

import (
	"context"
	"net"
	"os/signal"
	"syscall"

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
	jion := flag.String("jion", "", "host:port of leader to join")
	flag.Parse()
	if *jion != "" {
		conn, err := grpc.Dial(*jion, grpc.WithInsecure())
		if err != nil {
			logs.Error(err)
			return
		}
		_, err = msg.NewKvClient(conn).Man(context.TODO(), &msg.Manreq{Op: 2, Value: []byte(config.Ohmkvcfg.Raft.Addr + ":" + config.Ohmkvcfg.Raft.Port)})
		if err != nil {
			logs.Error(err)
			return
		}
	}
	cluster := cluster.NewCluster()
	go cluster.Sync()
	lis, err := net.Listen("tcp", config.Ohmkvcfg.Ohmkv.Addr+":"+config.Ohmkvcfg.Ohmkv.Port)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	s := grpc.NewServer()
	msg.RegisterKvServer(s, &rpc.Service{Cluster: cluster})
	go s.Serve(lis)
	sigch := make(chan os.Signal)
	signal.Notify(sigch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	<-sigch
	cluster.Close()
}
