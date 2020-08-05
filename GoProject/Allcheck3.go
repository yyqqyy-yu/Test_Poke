package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)
var p [13]string={"2","3","4","5","6","7","8","9","T","J","Q","K","A"}
var shu [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var lspk []int
var lspk1 []int
var huase1 string
var huase2 string
var sz []int
var sz1 APZ
var ae APZ
var bb APZ
var ans []APZ
var jilu []string
var qkjl []string
var js string
var js1 string
var js2 string
var js3 string
var cal int
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
func qrth(anyp string ) (res string){

	for hs:=3;hs<10;hs=hs+2{
		if string(anyp[1])!=string(anyp[hs]){
			return "no"
			break
		}
	}
	return string(anyp[1])
}


func qrsz (filename string,poke string )(aa APZ){
	var js string
	var ae APZ
	var o int
	var jilu []string
	inputFile, inputError := os.Open(filename)
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
	//po:=
	json.Unmarshal([]byte(js), &ans)

	for i := 0; i < 5; i++ {
		ae.PZ[i] = string(poke[o])
		o = o + 2
	}
	for jl := 0; jl < 5; jl++ {
		jilu = append(jilu, ae.PZ[jl])
	}
	calcu = 0
	for js := 0; js < len(ans); js++ {
		calcu = 0
		for ls := 0; ls < 5; ls++ {
			for ls1 := 0; ls1 < 5; ls1++ {
				if ans[js].PZ[ls] == ae.PZ[ls1]{
					ae.PZ[ls1] = "isok"
					calcu = calcu + 1
					//fmt.Println(calcu)
					break
				}
			}}
		for ix, vl := range jilu {
			ae.PZ[ix]= vl
		}
		if calcu == 5 {
			ae = ans[js]
			//fmt.Println("chu xian =0")
			break
		}

	}
	return ae
}
func main(){
	inputF, inputE := os.Open("match2.json")
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
	for por, v := range answe.Matches {
		huase1=qrth(v.Alice)
		huase2=qrth(v.Bob)
		if huase1!="no"&&huase2!="no"{
			ae=qrsz("Allcheck1.json",v.Alice)
			bb=qrsz("Allcheck1.json",v.Bob)
			if ae.ZL==5&&bb.ZL==5{
				if ae.PZS[0]!=bb.PZS[0]{
					if ae.PZS[0]>bb.PZS[0]{
						answe.Matches[por].Result=1
						goto ss
					}else {
						answe.Matches[por].Result=2
						goto ss
					}
				}else {
					answe.Matches[por].Result=3
					goto ss
				}
			}else if ae.ZL==5{
				answe.Matches[por].Result=1
				goto ss
			}else if bb.ZL==5{
				answe.Matches[por].Result=2
				goto ss
			}else {
				for z:=12;z<13;z--{
					for x:=0;x<10;x=x+2{
						if p[z]==string(v.Alice[x]){
							lspk=append(lspk, shu[z])
						}
						if p[z]==string(v.Bob[x]){
							lspk1=append(lspk1,shu[z])
						}
					}

				}
				if lspk[0]!=lspk1[0]{
					if lspk[0]>lspk1[0]{
						answe.Matches[por].Result=1
						goto ss
					}else {
						answe.Matches[por].Result=2
						goto ss
					}

				}else {
					if lspk[1]!=lspk1[1]{
						if lspk[1]>lspk1[1]{
							answe.Matches[por].Result=1
							goto ss
						}else {
							answe.Matches[por].Result=2
							goto ss
						}

					}else {
						if lspk[2]!=lspk1[2]{
							if lspk[2]>lspk1[2]{
								answe.Matches[por].Result=1
								goto ss
							}else {
								answe.Matches[por].Result=2
								goto ss
							}

						}else {
							if lspk[3]!=lspk1[3]{
								if lspk[3]>lspk1[3]{
									answe.Matches[por].Result=1
									goto ss
								}else {
									answe.Matches[por].Result=2
									goto ss
								}

							}else {
								if lspk[4]!=lspk1[4]{
									if lspk[4]>lspk1[4]{
										answe.Matches[por].Result=1
										goto ss
									}else {
										answe.Matches[por].Result=2
										goto ss
									}

								}else {
									answe.Matches[por].Result=3
									goto ss
								}
							}
						}
					}
				}

			}
		}else if huase1!="no"&&huase2=="no"{
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
			//po:=
			json.Unmarshal([]byte(js), &ans)
			o = 0
			for i := 0; i < 5; i++ {
				ae.PZ[i] = string(v.Alice[o])
				o = o + 2
			}
			for jl := 0; jl < 5; jl++ {
				jilu = append(jilu, ae.PZ[jl])
			}
			calcu = 0
			for js := 0; js < len(ans); js++ {
				calcu = 0
				for ls := 0; ls < 5; ls++ {
					for ls1 := 0; ls1 < 5; ls1++ {
						if ans[js].PZ[ls] == ae.PZ[ls1]{
							ae.PZ[ls1] = "isok"
							calcu = calcu + 1
							//fmt.Println(calcu)
							break
						}
					}}
				for ix, vl := range jilu {
					ae.PZ[ix]= vl
				}
				if calcu == 5 {
					ae = ans[js]
					//fmt.Println("chu xian =0")
					break
				}

			}
			if ae.ZL==5{
				answe.Matches[por].Result=1
				goto ss
			}else {
				var cal2 int
				for i:=0;i<5;i=i+2{
					for p:=0;p<10;p=p+2{
						if string(v.Bob[i])==string(v.Bob[p]){
							cal2=cal2+1
						}
					}
					cal2=cal2-1
					if cal2==3{
						answe.Matches[por].Result=2
						goto ss
					}else if cal2==2{
						var cal1 int
						for u:=0;u<7;u=u+2{
							for k:=0;k<10;k=k+2{
								if string(v.Bob[u])==string(v.Bob[k]){
									cal1=cal1+1
								}
							}
							cal1=cal1-1
							if cal1==1{
								answe.Matches[por].Result=2
								goto ss
							}else{
								answe.Matches[por].Result=1
								goto ss
							}
							cal1=0
						}
					}
					cal2=0
				}

			}
			t4:=time.Now()
			fmt.Println(t4.Sub(t3))
			fmt.Println(test1)
		}else if huase1=="no" &&huase2!="no"{
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
				js1 = js1 + inputString
			}
			//fmt.Printf("all: \n%s", js)
			//json字符串转化为结构体
			//po:=
			json.Unmarshal([]byte(js1), &ans)
			o = 0
			for i := 0; i < 5; i++ {
				ae.PZ[i] = string(v.Bob[o])
				o = o + 2
			}
			for jl := 0; jl < 5; jl++ {
				jilu = append(jilu, ae.PZ[jl])
			}
			calcu = 0
			for js := 0; js < len(ans); js++ {
				calcu = 0
				for ls := 0; ls < 5; ls++ {
					for ls1 := 0; ls1 < 5; ls1++ {
						if ans[js].PZ[ls] == ae.PZ[ls1]{
							ae.PZ[ls1] = "isok"
							calcu = calcu + 1
							//fmt.Println(calcu)
							break
						}
					}}
				for ix, vl := range jilu {
					ae.PZ[ix]= vl
				}
				if calcu == 5 {
					ae = ans[js]
					//fmt.Println("chu xian =0")
					break
				}
			}
			if ae.ZL==5{
				answe.Matches[por].Result=2
				goto ss
			}else {
				var cal2 int
				for i:=0;i<5;i=i+2{
					for p:=0;p<10;p=p+2{
						if string(v.Alice[i])==string(v.Alice[p]){
							cal2=cal2+1
						}
					}
					cal2=cal2-1
					if cal2==3{
						answe.Matches[por].Result=1
						goto ss
					}else if cal2==2{
						var cal1 int
						for u:=0;u<7;u=u+2{
							for k:=0;k<10;k=k+2{
								if string(v.Alice[u])==string(v.Alice[k]){
									cal1=cal1+1
								}
							}
							cal1=cal1-1
							if cal1==1{
								answe.Matches[por].Result=1
								goto ss
							}else{
								answe.Matches[por].Result=2
								goto ss
							}
							cal1=0
						}
					}
					cal2=0
				}

			}
			t4:=time.Now()
			fmt.Println(t4.Sub(t3))
			fmt.Println(test1)
		}else {

			for i:=0;i<5;i=i+2{
				for p:=0;p<10;p=p+2{
					if string(v.Bob[i])==string(v.Bob[p]){
						cal2=cal2+1
					}
				}
				cal2=cal2-1
				if cal2==3{
					answe.Matches[por].Result=2
					goto ss
				}else if cal2==2{
					var cal1 int
					for u:=0;u<7;u=u+2{
						for k:=0;k<10;k=k+2{
							if string(v.Bob[u])==string(v.Bob[k]){
								cal1=cal1+1
							}
						}
						cal1=cal1-1
						if cal1==1{
							answe.Matches[por].Result=2
							goto ss
						}else{
							answe.Matches[por].Result=1
							goto ss
						}
						cal1=0
					}
				}
				cal2=0
			}
		}

		ss:jilu = qkjl
	}
}
