// Package main

/*
课后练习 2.2
------------

编写一个 HTTP 服务器

1. 接收客户端 request，将 request 中带的 header 写入 response header
2. 读取当前系统的环境变量中的 VERSION 配置，写入 response header
3. Server 端记录访问日至包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 荡访问 localhost/healthz 时，应返回 200

思考：

1. 加入优雅终止的处理？
2. 工程化的思考？文档？测试？
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func healthz(writer http.ResponseWriter, request *http.Request) {
	// 获取环境变量
	version := os.Getenv("VERSION")
	//fmt.Println("VERSION:", version)
	writer.Header().Set("VERSION", version)
	// 遍历所有请求头，并加入到响应头中
	//fmt.Println("[request.Header]")
	for key := range request.Header {
		writer.Header().Del(key)
		values := request.Header[key]
		//fmt.Println(key, ":", values)
		for i := range values {
			writer.Header().Add(key, values[i])
		}
	}
	statusCode := 200
	writer.WriteHeader(statusCode)

	// 打印日志
	// TODO: 不知道什么原因 glog 无法配置输出到控制台中，即使运行参数指定了 -logtostderr=true
	//glog.V(4).Info(
	//	time.Now().Format("2006-01-02 15:04:05.000"),
	//	request.URL,
	//	fmt.Sprintf("[%d]", statusCode),
	//	request.RemoteAddr,
	//)

	fmt.Println(
		time.Now().Format("2006-01-02 15:04:05.000"),
		request.URL,
		fmt.Sprintf("[%d]", statusCode),
		request.RemoteAddr,
	)

	_, _ = fmt.Fprintln(writer, "Hello world!")
}

func main() {
	//defer glog.Flush()
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
