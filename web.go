package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func logWrite(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	gLog(req.FormValue("msg"))
}

func readLogs(w http.ResponseWriter, req *http.Request) {
	if len(_projID) > 0 && len(_logName) > 0 {
		logReader(_projID, _logName, w)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Println("hello")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func bucketEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("bucket")
	if r.Method == "POST" {
		for k, v := range r.Header {
			if strings.HasPrefix(k, "X-Goog") {
				msg := k + " IS " + strings.Join(v, "")
				fmt.Println(msg)
				gLog(msg)
			}
		}
	}
}

func webServer(port int) {
	checkCert()
	addr := fmt.Sprintf(":%d", port)
	s := &http.Server{
		Addr: addr,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/log/read", readLogs)
	http.HandleFunc("/log/write", logWrite)
	http.HandleFunc("/bucket/event", bucketEvent)
	log.Println("Start server --", addr)
	gLog("Start server -- " + addr)
	err := s.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func readPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		return string(body), err
	}
	return "", err
}
