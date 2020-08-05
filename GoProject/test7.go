package main


import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"bytes"
	"flag"
	"os"
	"strconv"
	"strings"
	"time"
)


type APZ struct {
	PZ [5]string
	PZS [5]int
	DX int
	ZL int
}
type FLPZ struct {
	FPZ []APZ
}

type HJJ struct {
	Alice string `json:"alice"`
	Bob    string `json:"bob"`
	Result    int `json:"result"`
}
type APP struct{
	Matches  []HJJ `json:"matches"`
}
func main(){

var (
	err error
	data []byte
	answer APP

)
//从配置文件读取转换为struct
configfile := flag.String("configfile", "/usr/match.json", "config file")
flag.Parse()
data, err = ioutil.ReadFile(*configfile)
if err != nil {
fmt.Println(err)
return
}
data = bytes.TrimPrefix(data, []byte("\xef\xbb\xbf"))

err = json.Unmarshal(data, &answer)
if err != nil {
fmt.Println(err)
return
} else {
//or _,v:=range answer.Matches{
	//p:=v.Alice[0:3]
	//fmt.Println(p)
//}
	//s:="As10s10sQsJs"
	l:="A2345678910JQK"
	fmt.Println(strings.IndexAny(l,"K"))
	var oo [5]string=[5]string{"a","1","2","3","4"}
	var pp [5]string=[5]string{"a","1","2","3","4"}
	var calcu int
	fmt.Println(len(oo))
	for ls:=0;ls<5;ls++{

		for ls1:=0;ls1<5;ls1++{

			if pp[ls]==oo[ls1] {

				calcu=calcu+1
				break

			}

		}


	}
	if (0-1)>=0 {
		fmt.Println("fuck you you")
	}
	var t []int
	for i:=0;i<5;i++{
		t=append(t, i)
	}
	fmt.Println(calcu)
	fmt.Println(t[1])
	b,_:=strconv.Atoi("222222")

	fmt.Println(b)
	fmt.Println()
	t1:=time.Now()
	inputF, inputE := os.Open("match1.json")
	if inputE != nil {
		return
	}
	defer inputF.Close()

	//按行读取文件
	var kkk string

	inputR := bufio.NewReader(inputF)
	for {
		inputString, readerError := inputR.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		kkk= kkk + inputString
	}
	//fmt.Println(kkk)


	//json字符串转化为结构体
	fmt.Println("-------------------------------------------------")
	var answe APP

	json.Unmarshal([]byte(kkk),&answe)
	t2:=time.Now()
var o int
	fmt.Println(t2.Sub(t1))
	var test1 []APZ
	var ae APZ
	var bb APZ
	var js string
	t3:=time.Now()
	inputFile, inputError := os.Open("AllCheck1.json")
	if inputError != nil {
		return
	}
	defer inputFile.Close()

	//按行读取文件

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		js = js + inputString
	}

	//fmt.Printf("all: \n%s", js)

	//json字符串转化为结构体
	fmt.Println("-------------------------------------------------")
	var ans []APZ
	var jilu []string
	var jilu1 []string
	//po:=
	json.Unmarshal([]byte(js), &ans)

	o = 0
	for i := 0; i < 5; i++ {
		ae.PZ[i] = string(answe.Matches[0].Alice[o])
		o = o + 2
	}
	test1 = append(test1, ae)

	o = 0
	for i := 0; i < 5; i++ {
		bb.PZ[i] = string(answe.Matches[0].Bob[o])
		o = o + 2
	}
	test1 = append(test1, bb)
	for jl := 0; jl < 5; jl++ {
		jilu = append(jilu, test1[0].PZ[jl])
		jilu1 = append(jilu1, test1[1].PZ[jl])
	}
	calcu = 0
	for js := 0; js < len(ans); js++ {
		calcu = 0
		for ls := 0; ls < 5; ls++ {
			for ls1 := 0; ls1 < 5; ls1++ {
				if ans[js].PZ[ls] == test1[0].PZ[ls1]{
					test1[0].PZ[ls1] = "isok"
					calcu = calcu + 1
					//fmt.Println(calcu)
					break
				}
			}}
		for ix, vl := range jilu {
			test1[0].PZ[ix] = vl
		}
		if calcu == 5 {
			test1[0] = ans[js]
			//fmt.Println("chu xian =0")
			break
		}
	}
	calcu = 0
	for js := 0; js < len(ans); js++ {
		calcu = 0
		for ls := 0; ls < 5; ls++ {
			for ls1 := 0; ls1 < 5; ls1++ {
				if ans[js].PZ[ls] == test1[1].PZ[ls1] {
					test1[1].PZ[ls1] = "isok"
					calcu = calcu + 1
					break
				}
			}
		}
		for ix, vl := range jilu1 {
			test1[1].PZ[ix] = vl
		}
		if calcu == 5 {
			test1[1] = ans[js]
			//fmt.Println("chu xian =1")
			break
		}
	}
	t4:=time.Now()
	fmt.Println(t4.Sub(t3))
	fmt.Println(test1)
}
}