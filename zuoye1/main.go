package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {

	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key:%s,Heaher value:%s \n", k, v)
			w.Header().Set(k, vv)
		}

	}

	clientip := getCurrentIP(r)
	log.Printf("clientIP:%s\n", clientip)
	log.Printf("client response code: %v", 200)

}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-real-Ip")
	if ip == "" {
		//fmt.Println(r.RemoteAddr)
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")

}
func main() {
	mux := http.NewServeMux()
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start httpd failed, err:%s\n", err.Error())

	}
}
