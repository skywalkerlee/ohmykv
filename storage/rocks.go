package storage

import (
	"github.com/astaxie/beego/logs"
	"github.com/skywalkerlee/ohmykv/config"
	"github.com/tecbot/gorocksdb"
)

type Iterm struct {
	Err   error
	Key   []byte
	Value []byte
}

type RocksStorage struct {
	db *gorocksdb.DB
}

var rs *RocksStorage

func GetRS() *RocksStorage {
	return rs
}

func NewRocksStorage() *RocksStorage {
	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
	opts := gorocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, config.Ohmkvcfg.Raft.StorageBackendPath)
	if err != nil {
		logs.Info(err)
	}
	rs = &RocksStorage{}
	rs.db = db
	return rs
}

func (rs *RocksStorage) Put(key []byte, value []byte) error {
	wo := gorocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()
	return rs.db.Put(wo, key, value)
}

func (rs *RocksStorage) Get(key []byte) ([]byte, error) {
	ro := gorocksdb.NewDefaultReadOptions()
	defer ro.Destroy()
	slice, err := rs.db.Get(ro, key)
	return slice.Data(), err
}

func (rs *RocksStorage) Del(key []byte) error {
	wo := gorocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()
	return rs.db.Delete(wo, key)
}

func (rs *RocksStorage) Close() {
	if rs.db != nil {
		rs.db.Close()
	}
}

func (rs *RocksStorage) Recv(ch chan *Iterm, seek ...[]byte) {
	go rs.getall(ch)
}

func (rs *RocksStorage) getall(ch chan *Iterm, seek ...[]byte) {
	ro := gorocksdb.NewDefaultReadOptions()
	defer ro.Destroy()
	ro.SetFillCache(false)
	it := rs.db.NewIterator(ro)
	defer it.Close()
	if len(seek) > 0 {
		it.Seek(seek[0])
	} else {
		it.SeekToFirst()
	}
	for ; it.Valid(); it.Next() {
		if err := it.Err(); err != nil {
			ch <- &Iterm{Err: err}
			break
		}
		key := make([]byte, len(it.Key().Data()))
		value := make([]byte, len(it.Key().Data()))
		copy(key, it.Key().Data())
		copy(value, it.Value().Data())
		ch <- &Iterm{Err: nil, Key: key, Value: value}
	}
	close(ch)
}
