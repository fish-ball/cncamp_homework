// Package main
/**
课后练习 1.1
------------

编写一个小程序：

给定一个字符串数组: [“I”,“am”,“stupid”,“and”,“weak”]

用 for 循环遍历该数组并修改为: [“I”,“am”,“smart”,“and”,“strong”]
*/
package main

import "fmt"

func main() {
	arr := []string{"I", "am", "stupid", "and", "weak"}
	arr[2] = "smart"
	arr[4] = "strong"
	fmt.Println(arr)
}
