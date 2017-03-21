package main

import (
	"io"
	"os"

	"math/rand"
	"time"

	"bufio"

	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/msg"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	addr := []string{"127.0.0.1:33881", "127.0.0.1:33882", "127.0.0.1:33883"}
	conn, err := grpc.Dial(addr[rand.New(rand.NewSource(time.Now().Unix())).Intn(3)], grpc.WithInsecure())
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	client := msg.NewKvClient(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			logs.Error(err)
			os.Exit(1)
		}
		word := strings.Split(line, " ")
		if line == "iter\n" {
			rc, err := client.Iterator(context.Background(), &msg.Iterreq{Op: 1})
			if err == io.EOF {
				return
			}

			for {
				resp, err := rc.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					logs.Error(err)
					return
				}
				logs.Info(resp)
			}
		}
		switch word[0] {
		case "get":
			resp, err := client.Read(context.Background(), &msg.Readreq{Key: []byte(word[1][:len(word[1])-1])})
			if err != nil {
				logs.Error(err)
				os.Exit(1)
			}
			logs.Info(resp.String())
		case "put":
			resp, err := client.Write(context.Background(), &msg.Writereq{Op: 1, Key: []byte(word[1]), Value: []byte(word[2][:len(word[2])-1])})
			if err != nil {
				logs.Error(err)
				os.Exit(1)
			}
			logs.Info(resp.String())
		case "del":
			resp, err := client.Write(context.Background(), &msg.Writereq{Op: 2, Key: []byte(word[1][:len(word[1])-1])})
			if err != nil {
				logs.Error(err)
				os.Exit(1)
			}
			logs.Info(resp.String())
		case "iter":
			var rc msg.Kv_IteratorClient
			if len(word) > 1 {
				rc, err = client.Iterator(context.Background(), &msg.Iterreq{Op: 2, Seek: []byte(word[1][:len(word[1])-1])})
				if err == io.EOF {
					return
				}
			} else {
				rc, err = client.Iterator(context.Background(), &msg.Iterreq{Op: 1})
				if err == io.EOF {
					return
				}
			}
			for {
				resp, err := rc.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					logs.Error(err)
					return
				}
				logs.Info(resp)
			}
		}

	}
}
