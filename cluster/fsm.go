package cluster

import (
	"encoding/binary"
	"io"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/raft"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/skywalkerlee/ohmykv/msg"
	"github.com/skywalkerlee/ohmykv/storage"
)

type storageFSM struct {
	storage *storage.RocksStorage
	lock    *sync.Mutex
}

func newStorageFSM(storage *storage.RocksStorage) *storageFSM {
	return &storageFSM{
		storage: storage,
		lock:    &sync.Mutex{},
	}
}

func (fsm *storageFSM) Apply(log *raft.Log) interface{} {
	fsm.lock.Lock()
	defer fsm.lock.Unlock()
	msg := &msg.Req{}
	err := proto.Unmarshal(log.Data, msg)
	if err != nil {
		logs.Error(err)
		return nil
	}
	if msg.GetOp() == 1 {
		fsm.storage.Put(msg.GetKey(), msg.GetValue())
	} else if msg.GetOp() == 2 {
		fsm.storage.Del(msg.GetKey())
	} else {
		config.Leader.Addr = string(msg.GetKey()) + ":" + string(msg.GetValue())
	}
	return "Apply Successful"
}

func (fsm *storageFSM) Snapshot() (raft.FSMSnapshot, error) {
	fsm.lock.Lock()
	defer fsm.lock.Unlock()
	return &StorageSnapshot{fsm.storage}, nil
}

func (s *storageFSM) Restore(inp io.ReadCloser) error {
	logs.Debug("Restore")
	s.lock.Lock()
	defer s.lock.Unlock()
	size := make([]byte, 2)
	for {
		_, err := inp.Read(size)
		if err == io.EOF {
			break
		}
		if err != nil {
			logs.Error(err)
			return err
		}
		tmp := make([]byte, int(binary.BigEndian.Uint16(size)))
		_, err = inp.Read(tmp)
		if err != nil {
			logs.Error(err)
			return err
		}
		msgDecode := &msg.Dump{}
		if err = proto.Unmarshal(tmp, msgDecode); err != nil {
			logs.Error(err)
			return err
		}
		if err = s.storage.Put(msgDecode.GetKey(), msgDecode.GetValue()); err != nil {
			logs.Error(err)
			return err
		}
	}
	return nil
}

type StorageSnapshot struct {
	storage *storage.RocksStorage
}

func (s *StorageSnapshot) Persist(sink raft.SnapshotSink) error {
	defer sink.Close()
	ch := make(chan *storage.Iterm, 1000)
	s.storage.Recv(ch)
	for tmp := range ch {
		if tmp.Err != nil {
			logs.Error(tmp.Err)
			return tmp.Err
		}
		msg := &msg.Dump{
			Key:   tmp.Key,
			Value: tmp.Value,
		}
		msdEncode, err := proto.Marshal(msg)
		if err != nil {
			logs.Error(err)
			return err
		}
		bSize := uint16(len(msdEncode))
		buf := make([]byte, bSize+2)
		binary.BigEndian.PutUint16(buf[:2], bSize)
		copy(buf[2:], msdEncode)
		if _, err = sink.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

func (s *StorageSnapshot) Release() {
	logs.Debug("Release the Snapshot")
}
