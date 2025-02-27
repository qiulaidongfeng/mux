package mux

import (
	"fmt"
	"net/http"
)

func ExampleStd() {
	m := New()
	a, b := http.NewServeMux(), http.NewServeMux()
	a.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "网站1")
	})
	b.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "网站2")
	})
	m.AddStd("a.com", a)
	m.AddStd("b.com", b)
	/*
		s := http.Server{
			Addr:    ":443",
			Handler: m,
		}
		s.ListenAndServeTLS("./cert.pem", "./key.pem")
	*/

	// 这样，a.com和b.com是不同的域名，可以在一个vps使用不同的域名搭建多个网站
	// 通过域名，决定响应网站1还是网站2
	// 通过https://ip，访问的是a.com，响应的是网站1
}
