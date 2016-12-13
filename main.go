package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
)

var (
	_projID  = "udy-demo"
	_logName = "bucket-log"
	lggr     = &gLogger{}
	localIP  = myIP()
)

const version = 0.01

func login() {
	log.Println("logging in")
	creds(_projID)
	var err error
	lggr, err = newClient(_projID)
	if err != nil {
		panic(err)
	}
}

func gLog(msg string) {
	lggr.writeEntry(_logName, msg)
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for sig := range c {
			gLog("Got signal:" + sig.String())
			os.Exit(1)
		}
	}()
}

func main() {
	logs := false
	flag.BoolVar(&logs, "logs", false, "show logs and exit")
	flag.Parse()
	if logs {
		logReader(_projID, _logName, os.Stdout)
		return
	}
	options()
	login()
	webServer(443)
}
