package main

import (
	"fmt"
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/service"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}

	http.HandleFunc("/", service.IndexHandler)
	http.HandleFunc("/api/count", service.CounterHandler)
	// 静态资源：/res/xxx 映射到 res 文件夹下的对应文件
	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("res"))))

	log.Fatal(http.ListenAndServe(":80", nil))
}
