package conf

import (
	"flag"
)
import "github.com/jinzhu/configor"

var (
	confPath string
	Conf     *Config
)

func init() {
	flag.StringVar(&confPath, "conf", "worker-example.json", "default config path.")
}

type Config struct {
	EtcdEndpoints         []string `json:"etcd_endpoints"`
	EtcdDialTimeout       int      `json:"etcd_dial_timeout"`
	MongodbUri            string   `json:"mongodb_uri"`
	MongodbConnectTimeout int      `json:"mongodb_connect_timeout"`
	JobLogBatchSize       int      `json:"job_log_batch_size"`
	JobLogCommitTimeout   int      `json:"job_log_commit_timeout"`
}

func Init() (err error) {
	Conf = Default()
	//_, err = toml.DecodeFile(confPath, &Conf)
	err = configor.Load(&Conf, confPath)
	return
}
func Default() *Config {
	return &Config{

	}
}
