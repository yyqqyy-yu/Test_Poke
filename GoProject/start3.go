package main

import (
	"bufio"
	"encoding/json"

	"fmt"
	"io"
	"os"
)


type Matches struct {
	Alice string `json:"alice"`
	Bob    string `json:"bob"`
	Result    int `json:"result"`
}
type List struct{
	Matches  []Matches `json:"matches"`
}

func main() {
	//打开文件
	inputFile, inputError := os.Open("/usr/ma.json")
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
	var sss List
	json.Unmarshal([]byte(s),&sss)
	fmt.Println(sss)

}