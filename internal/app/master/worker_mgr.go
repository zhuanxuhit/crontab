package master

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/zhuanxuhit/crontab/internal/app/common"
	"github.com/zhuanxuhit/crontab/internal/app/worker/conf"
	"time"
)

// /cron/workers/
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
	// 建立连接
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

// 获取在线worker列表
func (workerMgr *WorkerMgr) ListWorkers() (workerArr []string, err error) {
	var (
		getResp *clientv3.GetResponse
		kv *mvccpb.KeyValue
		workerIP string
	)

	// 初始化数组
	workerArr = make([]string, 0)

	// 获取目录下所有Kv
	if getResp, err = workerMgr.kv.Get(context.TODO(), common.JobWorkerDir, clientv3.WithPrefix()); err != nil {
		return
	}

	// 解析每个节点的IP
	for _, kv = range getResp.Kvs {
		// kv.Key : /cron/workers/192.168.2.1
		workerIP = common.ExtractWorkerIP(string(kv.Key))
		workerArr = append(workerArr, workerIP)
	}
	return
}