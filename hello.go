package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
)

func init() {
	router.GET("/hello/:name", Hello)
}

// Hello func
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if log.Level == logrus.DebugLevel {
		dump, _ := httputil.DumpRequest(r, true)
		log.Debugf("%q", dump)
	}

	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
