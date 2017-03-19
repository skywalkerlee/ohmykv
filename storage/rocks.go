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
	return &RocksStorage{db}
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
	rs.db.Close()
}

func (rs *RocksStorage) Recv(ch chan *Iterm) {
	go rs.getall(ch)
}

func (rs *RocksStorage) getall(ch chan *Iterm) {
	ro := gorocksdb.NewDefaultReadOptions()
	defer ro.Destroy()
	ro.SetFillCache(false)
	it := rs.db.NewIterator(ro)
	defer it.Close()
	it.SeekToFirst()
	for ; it.Valid(); it.Next() {
		if err := it.Err(); err != nil {
			ch <- &Iterm{Err: err}
			continue
		}
		ch <- &Iterm{Err: nil, Key: it.Key().Data(), Value: it.Value().Data()}
	}
}
