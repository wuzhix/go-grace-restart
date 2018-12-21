package restart

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	store restore
)

func init() {
	//1、启动goroutine监听关闭信号
	go hookSignal()
	println("go hookSignal()")
	//2、获取上一个应用程序的pid
	lastPid := getPid()
	println("lastPid: ", lastPid)
	if lastPid != 0 {
		path := findFile(lastPid)
		println("lastPid path: ", path)
		//3、关闭上一个进程
		err := kill(lastPid)
		if err != nil {
			println(err.Error())
		} else {
			println("kill ", lastPid)
			for i := 0; i < 10; i++ {
				time.Sleep(time.Duration(1) * time.Second)
				if !exist(lastPid) {
					if store != nil {
						store.load()
					}
					err = deleteFile(path)
					if err != nil {
						println(err.Error())
					}
					println("deleteFile ", path)
					break
				}
			}
		}
	}
	savePid()
}

func hookSignal() {
	sig := make(chan os.Signal, 2)
	//监听关闭信号
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//阻塞等待关闭信号
	<-sig
	println("catch exit signal")
	//windows下无法捕获到process.Kill()...
	//保存现场
	if store != nil {
		store.save()
	}
	os.Exit(1)
}
