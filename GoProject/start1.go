package main

import (
	"bufio"
	"encoding/json"

	"fmt"
	"io"
	"os"
)


type Kind struct {
	Alice string `json:"alice"`
	Bob    string `json:"bob"`
	Result    int `json:"result"`
}
type All struct{
	Matches  []Kind `json:"matches"`
}

func main() {
	//打开文件
inputFile, inputError := os.Open("/usr/match.json")
	if inputError != nil {
		return
	}
	//defer inputFile.Close()

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
	var answer All
	json.Unmarshal([]byte(s),&answer)
	fmt.Println(answer)

}