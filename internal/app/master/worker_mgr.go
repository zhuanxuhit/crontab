package master

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/zhuanxuhit/crontab/internal/app/master/conf"
	"time"
)

type WorkerMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease

	config clientv3.Config
}

var workerMgr *WorkerMgr

func InitWorkerMgr() (err error) {
	var (
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
		config clientv3.Config
	)
	config = clientv3.Config{
		Endpoints:   conf.Conf.EtcdEndpoints,
		DialTimeout: time.Duration(conf.Conf.EtcdDialTimeout) * time.Millisecond,
	}
	if client, err = clientv3.New(config); err != nil {
		return
	}
	// 得到KV和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	workerMgr = &WorkerMgr{
		client: client,
		kv:     kv,
		lease:  lease,
		config: config,
	}
	return
}
