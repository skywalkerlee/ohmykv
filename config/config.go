package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/astaxie/beego/logs"
)

type config struct {
	Ohmkv struct {
		Addr   string `toml:"addr"`
		Port   string `toml:"port"`
		LogDir string `toml:"log_dir"`
	} `toml:"ohmkv"`
	Raft struct {
		Addr               string   `toml:"addr"`
		Port               string   `toml:"port"`
		Peers              []string `toml:"peers"`
		PeerStorage        string   `toml:"peer_storage"`
		SnapshotStorage    string   `toml:"snapshot_storage"`
		StorageBackendPath string   `toml:"storage_backend_path"`
		RaftLogPath        string   `toml:"raft_log_path"`
		ApplyLogPath       string   `toml:"apply_log_path"`
	} `toml:"raft"`
}

var Ohmkvcfg config

func init() {
	if _, err := toml.DecodeFile("./ohmkv.conf", &Ohmkvcfg); err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+Ohmkvcfg.Ohmkv.LogDir+`omq.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
