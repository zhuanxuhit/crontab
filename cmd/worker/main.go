package main

import (
	"flag"
	log "github.com/golang/glog"
	"github.com/zhuanxuhit/crontab/internal/app/worker"
	"github.com/zhuanxuhit/crontab/internal/app/worker/conf"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func InitEnv() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	flag.Parse()
	// 加载配置
	if err := conf.Init(); err != nil {
		panic(err)
	}
	InitEnv()

	// 服务注册
	if err = worker.InitRegister(); err != nil {
		panic(err)
	}
	// 启动日志协程
	if err = worker.InitLogSink(); err != nil {
		panic(err)
	}
	// 启动执行器
	if err = worker.InitExecutor(); err != nil {
		panic(err)
	}

	// 启动调度器
	if err = worker.InitScheduler(); err != nil {
		panic(err)
	}

	// 初始化任务管理器
	if err = worker.InitJobMgr(); err != nil {
		panic(err)
	}

	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("crontab-master get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}

}
