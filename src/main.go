package main

import (
	"fmt"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/routers"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 信号处理
	dealSignal()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.GetGlobalConfig().ServerSetting.HttpPort),
		Handler:        routers.InitRouter(),
		ReadTimeout:    config.GetGlobalConfig().ServerSetting.ReadTimeout,
		WriteTimeout:   config.GetGlobalConfig().ServerSetting.WriteTimeout,
		MaxHeaderBytes: config.GetGlobalConfig().AppSetting.MaxHeaderBytes,
	}

	log.Infof("ready start http server listening %s", config.GetGlobalConfig().ServerSetting.HttpPort)

	err := server.ListenAndServe()
	if err != nil {
		log.Warnf("server init failed, err:%s", err)
	}

	log.Infof("start http server listening %s", config.GetGlobalConfig().ServerSetting.HttpPort)
}

func dealSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		for s := range sigs {
			switch s {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT:
				log.Warnf("got signal:%v and try to exit: ", s)
				os.Exit(0)
			default:
				log.Warnf("other signal:%v: ", s)
			}
		}
	}()
}
