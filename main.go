package main

import (
	"flag"
	"fmt"
	"github.com/domeos/sender/cron"
	"github.com/domeos/sender/g"
	"github.com/domeos/sender/http"
	"github.com/domeos/sender/redis"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *version {
		fmt.Println("falcon-sender: ", g.FALCON_SENDER_VERSION)
		fmt.Println("domeos-sender: ", g.DOMEOS_VERSION)
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	g.ParseConfig(*cfg)
	cron.InitWorker()
	redis.InitConnPool()
	g.InitDatabase()

	go cron.UpdateApiConfig()
	go http.Start()
	go cron.ConsumeSms()
	go cron.ConsumeMail()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Println()
		redis.ConnPool.Close()
		os.Exit(0)
	}()

	select {}
}