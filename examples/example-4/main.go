package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
    // we need a webserver to get the pprof webserver
	router := mux.NewRouter()
	router.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	router.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	router.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	router.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	router.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	router.Handle("/debug/pprof/{cmd}", http.HandlerFunc(pprof.Index)) // special handling for Gorilla mux
	
	err := http.ListenAndServe("localhost:6060", router)
	fmt.Println("pprof server listen failed: %v", err)
}