package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:2002" // 自己监听的地址

func main() {
	//127.0.0.1:2002/xxx
	//127.0.0.1:2003/base/xxx              请求 127.0.0.1:2002/xxx ， 会被转发到 127.0.0.1:2003/base/xxx
	rs1 := "http://127.0.0.1:2003/base"  // 需要转发到哪里   目标地址
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	proxy := httputil.NewSingleHostReverseProxy(url1)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}