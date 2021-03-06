package main

import (
	"flag"
	log "github.com/golang/glog"
	"github.com/zhuanxuhit/crontab/internal/app/master"
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
	if err := conf.Init(); err != nil {
		panic(err)
	}
	InitEnv()

	// 初始化服务发现模块
	if err = master.InitWorkerMgr(); err != nil {
		panic(err)
	}

	// 日志管理器
	if err = master.InitLogMgr(); err != nil {
		panic(err)
	}

	//  任务管理器
	if err = master.InitJobMgr(); err != nil {
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
