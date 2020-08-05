package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonInfo struct {
	alice    string `json:"alice"`
	bob     string `json:"bob"`
	result  int `json:"result"`
}

func main() {

	readFile()
}

func readFile() {

	filePtr, err := os.Open("/usr/local/match.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var o interface{}

	//2.声明结构体

	var b []byte = make([]byte, 1024*1024*1024)
	var i int;
breakhere:
	i,_=filePtr.Read(b);
	if i>0 {
		err := json.Unmarshal(b, &o)
		// 向file的文件中读取b大小的字节
		if err != nil {
			fmt.Println("json err:", err)
		}

		goto breakhere }
	//3.json解析到结构体


	fmt.Println("s",o)



}
