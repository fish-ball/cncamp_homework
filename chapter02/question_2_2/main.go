// Package main

/*
课后练习 2.2
------------

编写一个 HTTP 服务器

1. 接收客户端 request，将 request 中带的 header 写入 response header
2. 读取荡起那系统的环境变量中的 VERSION 配置，写入 response header
3. Server 端记录访问日至包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4. 荡访问 localhost/healthz 时，应返回 200

思考：

1. 加入优雅终止的处理？
2. 工程化的思考？文档？测试？
*/
package main

import (
	"fmt"
	"github.com/golang/glog"
)

func main() {
	fmt.Println()
	glog.V(2).Info("Hello go")

}
