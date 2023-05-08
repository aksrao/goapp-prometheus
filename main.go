package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Device struct {
	Id       int    `json:"id"`
	Mac_addr string `json:"mac-addr"`
	OS       string `json:"os"`
}

var dvs []Device

func init() {
	dvs = []Device{
		{1, "a9:bc:63:5d:75:dd", "Windows"},
		{2, "60:7d:a3:7f:d7:0a", "IOS"},
		{3, "42:0c:41:7d:25:58", "Android"},
		{4, "c8:3f:b7:f5:e1:0a", "MacOS"},
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome,%q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You are in home page")
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
