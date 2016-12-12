package main

import "os"

var (
	_projID  = "udy-demo"
	_logName = "bucket-log"
	lggr     = &gLogger{}
)

func init() {
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

func hello() {
	bucketName := _projID + ".appspot.com"
	readOut(bucketName, "hello.txt")
}

func sayWhat() {
	gLog("hey man")
	logReader(_projID, _logName, os.Stdout)
}

func main() {
	/*
		sayWhat()
		return
		hello()
	*/
	webServer(443)
}
