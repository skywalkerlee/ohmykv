package main

import (
	"os"

	"math/rand"
	"time"

	"io"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/msg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	addr := []string{"127.0.0.1:33881", "127.0.0.1:33882", "127.0.0.1:33883"}
	conn, err := grpc.Dial(addr[rand.New(rand.NewSource(time.Now().Unix())).Intn(3)], grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	client := msg.NewKvClient(conn)
	rc, err := client.Iterator(context.Background(), &msg.Iterreq{Op: 1})
	if err != nil {
		logs.Error(err)
		return
	}
	for {
		resp, err := rc.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Info(resp)
	}
}
