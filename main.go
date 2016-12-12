package main

import "log"

var (
	_projID  = "udy-demo"
	_logName = "bucket-log"
	lggr     = &gLogger{}
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

func main() {
	options()
	login()
	webServer(443)
}
