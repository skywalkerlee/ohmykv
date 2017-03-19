package cluster

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gogo/protobuf/proto"
	"github.com/hashicorp/raft"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/skywalkerlee/ohmykv/msg"
	"github.com/skywalkerlee/ohmykv/storage"
)

type Cluster struct {
	Leaderaddr string
	Raft       *raft.Raft
}

func NewCluster() *Cluster {
	cfg := raft.DefaultConfig()
	memStore := raft.NewInmemStore()
	if err := os.RemoveAll(config.Ohmkvcfg.Raft.StorageBackendPath); err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	rs := storage.NewRocksStorage()
	fsm := newStorageFSM(rs)
	snap, err := raft.NewFileSnapshotStore(config.Ohmkvcfg.Raft.SnapshotStorage, 3, nil)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	tranAddr := fmt.Sprintf("%s:%s", config.Ohmkvcfg.Raft.Addr, config.Ohmkvcfg.Raft.Port)
	tran, err := raft.NewTCPTransport(tranAddr, nil, 3, 2*time.Second, nil)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	peerStorage := raft.NewJSONPeers(config.Ohmkvcfg.Raft.PeerStorage, tran)
	ps := config.Ohmkvcfg.Raft.Peers
	logs.Debug(ps)
	peerStorage.SetPeers(ps)
	ps, err = peerStorage.Peers()
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	logs.Debug(ps)
	r, err := raft.NewRaft(cfg, fsm, memStore, memStore, snap, peerStorage, tran)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	go func() {
		for {
			logs.Debug(r.Leader())
			time.Sleep(time.Second)
		}
	}()
	return &Cluster{Raft: r}
}

func (cluster *Cluster) Sync() {
	for {
		time.Sleep(time.Second * 3)
		if cluster.Raft.Leader() == config.Ohmkvcfg.Raft.Addr+":"+config.Ohmkvcfg.Raft.Port {
			config.Leader.Addr = config.Ohmkvcfg.Raft.Addr + ":" + config.Ohmkvcfg.Ohmkv.Port
			msg := &msg.Req{
				Op:    3,
				Key:   []byte(config.Ohmkvcfg.Raft.Addr),
				Value: []byte(config.Ohmkvcfg.Ohmkv.Port),
			}
			msgencode, err := proto.Marshal(msg)
			if err != nil {
				logs.Error(err)
			}
			cluster.Raft.Apply(msgencode, time.Second)
		}
	}
}
