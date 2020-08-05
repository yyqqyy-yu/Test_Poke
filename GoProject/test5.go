package main
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type County struct{
	Code    int
	Name    string
}
type City struct {
	Children []County
	Code    int
	Name    string
}
type Province struct{
	Children []City
	Code    string
	Name    string
}

func main() {
	//打开文件
	inputFile, inputError := os.Open("/usr/local/something.json")
	if inputError != nil {
		return
	}
	defer inputFile.Close()

	//按行读取文件
	var s string
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		s = s + inputString
	}
	fmt.Printf("all: \n%s", s)

	//json字符串转化为结构体
	fmt.Println("-------------------------------------------------")
	var province []Province
	if err := json.Unmarshal([]byte(s), &province); err == nil {
		fmt.Println(province)
		for i:=0;i<len(province);i++ {
			fmt.Println(province[i].Name);
		}
	}else{
		fmt.Println("转换失败")
	}
}
