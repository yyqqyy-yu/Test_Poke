package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("～/Desktop/match.json")        // 读操作一般使用Open打开文件
	if err != nil {
		// 文件打开失败
		return
	}
	defer file.Close()    // 关文件
	var b []byte = make([]byte, 2*1024)
	var i int;
	breakhere:
	i,_=file.Read(b);
	if i>0 {

	// 向file的文件中读取b大小的字节
	var str string = string(b)
	fmt.Println(str)
	goto breakhere }
}