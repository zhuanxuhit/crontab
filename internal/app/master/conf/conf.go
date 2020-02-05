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
	flag.StringVar(&confPath, "conf", "master-example.json", "default config path.")
}

type Config struct {
	ApiPort               int      `json:"api_port"`
	ApiReadTimeout        int      `json:"api_read_timeout"`
	ApiWriteTimeout       int      `json:"api_write_timeout"`
	EtcdEndpoints         []string `json:"etcd_endpoints"`
	EtcdDialTimeout       int      `json:"etcd_dial_timeout"`
	WebRoot               string   `json:"web_root"`
	MongodbUri            string   `json:"mongodb_uri"`
	MongodbConnectTimeout int      `json:"mongodb_connect_timeout"`
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
