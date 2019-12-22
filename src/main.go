package main

import (
	"fmt"
	"capture/src/config"
	"flag"
	"os"
	"io"
	"capture/src/mylog"
)

func main()  {
	conf := new(config.Cfg)
	conf.Port = flag.String("port", "8080", "Listen port")
	conf.Raddr = flag.String("raddr", "", "Remote addr")
	conf.Log = flag.String("logFile", "", "log file path")
	conf.Monitor = flag.Bool("m", false, "monitor mode")
	conf.Tls = flag.Bool("tls", false, "tls connect")

	flag.Parse()

	var log io.WriteCloser
	var err error
	// init log
	if *conf.Log != "" {
		log, err = os.Create(*conf.Log)
		mylog.SetLog(log)
		if err != nil {
			mylog.Fatalln("fail to create log file " + err.Error())
		}
	} else {
		log = os.Stderr
		mylog.SetLog(log)
	}
	mylog.Fatalln("fail to create log file " + err.Error())

	fmt.Printf("aaa")
}
