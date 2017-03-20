package rpc

import (
	"time"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/gogo/protobuf/proto"
	"github.com/skywalkerlee/ohmykv/cluster"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/skywalkerlee/ohmykv/msg"
)

type Service struct {
	Cluster *cluster.Cluster
}

func (service *Service) Op(ctx context.Context, req *msg.Req) (*msg.Resp, error) {
	tmp, err := proto.Marshal(req)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	for i := 0; i < 3; i++ {
		if service.Cluster.Raft.Leader() != "" {
			if service.Cluster.Raft.Leader() == config.Ohmkvcfg.Raft.Addr+":"+config.Ohmkvcfg.Raft.Port {
				resp := service.Cluster.Raft.Apply(tmp, time.Millisecond*10)
				if err := resp.Error(); err != nil {
					logs.Error(err)
					return nil, err
				}
				if resp.Response() == "Apply Successful" {
					return &msg.Resp{Status: 1}, nil
				}
			}
			conn, err := grpc.Dial(config.Leader.Addr, grpc.WithInsecure())
			if err != nil {
				logs.Error(err)
				return nil, err
			}
			return msg.NewKvClient(conn).Op(context.TODO(), req)
		}
		time.Sleep(time.Millisecond * 10)
	}
	return nil, errors.New("cluster has no leader")
}

func (service *Service) Man(ctx context.Context, req *msg.Manreq) (*msg.Manresp, error) {
	switch req.GetOp() {
	case 1:
		return &msg.Manresp{Status: 1, Body: []byte(service.Cluster.Raft.Stats()["num_peers"])}, nil
	case 2:
		if err := service.Cluster.Raft.AddPeer(string(req.GetValue())).Error(); err != nil {
			return nil, err
		}
		return &msg.Manresp{Status: 1}, nil
	case 3:
		if err := service.Cluster.Raft.RemovePeer(string(req.GetValue())).Error(); err != nil {
			return nil, err
		}
		return &msg.Manresp{Status: 1}, nil
	default:
		return nil, errors.New("Unknown option code")
	}
}
