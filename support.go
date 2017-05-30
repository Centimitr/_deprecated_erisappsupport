package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var api API

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Write([]byte("Hello"))
	})
	http.HandleFunc("/book", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "text/json; charset=utf-8")
		req.ParseForm()
		locator := req.Form.Get("locator")
		keys := req.Form.Get("keys")

		api.GetBook(locator, keys, res)
	})
	http.HandleFunc("/book/page", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		locator := req.Form.Get("locator")
		keys := req.Form.Get("keys")
		page := req.Form.Get("page")

		api.GetBookPage(locator, keys, page, res)
	})
	http.HandleFunc("/pack", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "text/json; charset=utf-8")
		req.ParseForm()
		path := req.Form.Get("path")
		dst := req.Form.Get("dst")

		api.MakeErisFromFolder(path, dst)
	})

	// port config
	port := 3455
	if len(os.Args) > 1 {
		port, _ = strconv.Atoi(os.Args[1])
	}

	// http2 cert and key
	certFile := createTempFile("com.devbycm.eris.support", _TLS_CERT)
	keyFile := createTempFile("com.devbycm.eris.support", _TLS_KEY)
	defer os.Remove(certFile)
	defer os.Remove(keyFile)

	//log.Fatal(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", port), nil))
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf("127.0.0.1:%d", port), certFile, keyFile, nil))
}

//func mainx() {
//	server.HandleFunc("a", func() {
//		fmt.Println(1)
//	})
//	for {
//		var v string
//		fmt.Scan(&v)
//		server.Handle("a")
//	}
//}
