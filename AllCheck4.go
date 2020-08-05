package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)
type HJJ struct {
	Alice string `json:"alice"`
	Bob    string `json:"bob"`
	Result    int `json:"result"`
}
type APP struct{
	Matches  []HJJ `json:"matches"`
}
type LSPP struct{
	Pk [5]int
	Ddj int
}
var lspk []int
var lspk1 []int
var lspk3 []int
var lspk4 []int
var lslp LSPP
var lslp1 LSPP
var qk []int
var p [13]string=[13]string{"2","3","4","5","6","7","8","9","T","J","Q","K","A"}
var shu [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var dj1 int
var dj2 int
var huase1 string
var huase2 string
func qrth(anyp string ) (res string){

	for hs:=3;hs<10;hs=hs+2{
		if string(anyp[1])!=string(anyp[hs]){
			return "no"
			break
		}
	}
	return string(anyp[1])
}
func pddj (lspk []int)(lp LSPP){
	if lspk[0]==lspk[3]||lspk[1]==lspk[4]{
			if lspk[1]==lspk[4]{
				for i:=0;i<4;i++{
					lp.Pk[i]=lspk[i+1]
				}
				lp.Pk[4]=lspk[0]
				lp.Ddj=7
				return lp
			}else {
				for i:=0;i<5;i++{
					lp.Pk[i]=lspk[i]
				}
				lp.Ddj=7
				return lp
			}
	}else if lspk[0]==lspk[2]||lspk[1]==lspk[3]||lspk[2]==lspk[4]{
		if lspk[0]==lspk[2]{
			if lspk[3]==lspk[4]{
				for i:=0;i<5;i++{
					lp.Pk[i]=lspk[i]
				}
				lp.Ddj=6
				return lp
			}else {
				for i:=0;i<5;i++{
					lp.Pk[i]=lspk[i]
				}
				lp.Ddj=4
				return lp
			}
		}else if lspk[0]==lspk[1]{
			for i:=0;i<3;i++{
				lp.Pk[i]=lspk[i+2]
			}
			lp.Pk[3]=lspk[0]
			lp.Pk[4]=lspk[1]
			lp.Ddj=6
			return lp
		}else {
			if lspk[1]==lspk[3]{
				for i:=0;i<3;i++{
					lp.Pk[i]=lspk[i+1]
				}
				lp.Pk[3]=lspk[0]
				lp.Pk[4]=lspk[4]
				lp.Ddj=4
				return lp
			}else {
				for i:=0;i<3;i++{
					lp.Pk[i]=lspk[i+2]
				}
				lp.Pk[3]=lspk[0]
				lp.Pk[4]=lspk[1]
				lp.Ddj=4
				return lp
			}
		}
	}else if lspk[0]==lspk[1]||lspk[1]==lspk[2]||lspk[2]==lspk[3]||lspk[3]==lspk[4]{
		if lspk[0]==lspk[1]{
			if lspk[2]==lspk[3]{
				for i:=0;i<5;i++{
					lp.Pk[i]=lspk[i]
				}
				lp.Ddj=3
				return lp
			}else if lspk[3]==lspk[4]{
				for i:=0;i<2;i++{
					lp.Pk[i]=lspk[i]
				}
				for i:=2;i<4;i++{
					lp.Pk[i]=lspk[i+1]
				}
				lp.Pk[4]=lspk[2]
				lp.Ddj=3
				return lp
			}else {
				for i:=0;i<5;i++{
					lp.Pk[i]=lspk[i]
				}
				lp.Ddj=2
				return lp
			}
		}else if lspk[1]==lspk[2]{
			if lspk[3]==lspk[4]{
				for i:=0;i<2;i++{
					lp.Pk[i]=lspk[i+1]
				}
				for i:=2;i<4;i++{
					lp.Pk[i]=lspk[i+1]
				}
				lp.Pk[4]=lspk[0]
				lp.Ddj=3
				return lp
			}else {
				for i:=0;i<2;i++{
					lp.Pk[i]=lspk[i+1]
				}
				lp.Pk[2]=lspk[0]
				lp.Pk[3]=lspk[3]
				lp.Pk[4]=lspk[4]
				lp.Ddj=2
				return lp
			}
		}else if lspk[2]==lspk[3]{
				for i:=0;i<2;i++{
					lp.Pk[i]=lspk[i+2]
				}
				lp.Pk[2]=lspk[0]
				lp.Pk[3]=lspk[1]
				lp.Pk[4]=lspk[4]
				lp.Ddj=2
				return lp
			}else {
				for i:=0;i<2;i++{
					lp.Pk[i]=lspk[i+3]
				}
				lp.Pk[2]=lspk[0]
				lp.Pk[3]=lspk[1]
				lp.Pk[4]=lspk[2]
				lp.Ddj=2
				return lp
			}

	}else if lspk[0]-lspk[4]==4||lspk[0]==13&&lspk[1]==4{
		lp.Ddj=5
		return lp
	}else {
		lp.Ddj=1
		return lp
	}


}

func main(){
	t1:=time.Now()
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
	for por,v:=range answe.Matches{
		huase1=qrth(v.Alice)
		huase2=qrth(v.Bob)
		if huase1!="no"&&huase2!="no"{
			for z:=12;z>=0;z--{
				for x:=0;x<10;x=x+2{
					if p[z]==string(v.Alice[x]){
						lspk=append(lspk, shu[z])
					}
					if p[z]==string(v.Bob[x]){
						lspk1=append(lspk1,shu[z])
					}
				}
			}
			if (lspk[0]-lspk[4])==4||lspk[0]==13&&lspk[1]==4&&(lspk1[0]-lspk1[4])==4||lspk1[0]==13&&lspk1[1]==4{
				if lspk[0]==13&&lspk[1]==4&&lspk1[0]==13&&lspk1[1]==4{
					answe.Matches[por].Result=3
				}else if lspk[0]==13&&lspk[1]==4{
					answe.Matches[por].Result=2
				}else if lspk1[0]==13&&lspk1[1]==4{
					answe.Matches[por].Result=1
				}else if lspk[0]!=lspk1[0]{
					if lspk[0]>lspk1[0]{
						answe.Matches[por].Result=1
					}else {
						answe.Matches[por].Result=2
					}
				}else{
					answe.Matches[por].Result=3
				}
			}else if (lspk[0]-lspk[4])==4||lspk[0]==13&&lspk[1]==4{
				answe.Matches[por].Result=1
			}else if (lspk1[0]-lspk1[4])==4||lspk1[0]==13&&lspk1[1]==4{
				answe.Matches[por].Result=2
			}else {
				if lspk[0]!=lspk1[0]{
					if lspk[0]>lspk1[0]{
						answe.Matches[por].Result=1
					}else {
						answe.Matches[por].Result=2
					}
				}else {
					if lspk[1]!=lspk1[1]{
						if lspk[1]>lspk1[1]{
							answe.Matches[por].Result=1
						}else {
							answe.Matches[por].Result=2
						}
					}else {
						if lspk[2]!=lspk1[2]{
							if lspk[2]>lspk1[2]{
								answe.Matches[por].Result=1
							}else {
								answe.Matches[por].Result=2
							}
						}else {
							if lspk[3]!=lspk1[3]{
								if lspk[3]>lspk1[3]{
									answe.Matches[por].Result=1
								}else {
									answe.Matches[por].Result=2
								}
							}else {
								if lspk[4]!=lspk1[4]{
									if lspk[4]>lspk1[4]{
										answe.Matches[por].Result=1
									}else {
										answe.Matches[por].Result=2
									}
								}else {
									answe.Matches[por].Result=3
								}
							}
						}
					}
				}
			}
		}else if huase1!="no"{
			for z:=12;z>=0;z--{
				for x:=0;x<10;x=x+2{
					if p[z]==string(v.Alice[x]){
						lspk=append(lspk, shu[z])
					}
					if p[z]==string(v.Bob[x]){
						lspk1=append(lspk1,shu[z])
					}
				}
			}
			if lspk[0]==13&&lspk[1]==4{
				answe.Matches[por].Result=1
			}else if lspk[0]-lspk[4]==4{
				answe.Matches[por].Result=1
			}else {
				if lspk1[0]==lspk1[3]||lspk1[1]==lspk1[4]{
					answe.Matches[por].Result=2
				}else if lspk1[0]==lspk1[2]||lspk1[2]==lspk1[4]{
					if lspk1[0]==lspk1[2]{
						if lspk1[3]==lspk1[4]{
							answe.Matches[por].Result=2
						}else {
							answe.Matches[por].Result=1
						}
					}else if lspk1[2]==lspk1[4]{
						if lspk1[0]==lspk1[1]{
							answe.Matches[por].Result=2
						}else{
							answe.Matches[por].Result=1
						}
					}
				}else {
					answe.Matches[por].Result=1
				}
			}

		}else if huase2!="no"{
			for z:=12;z>=0;z--{
				for x:=0;x<10;x=x+2{
					if p[z]==string(v.Alice[x]){
						lspk=append(lspk, shu[z])
					}
					if p[z]==string(v.Bob[x]){
						lspk1=append(lspk1,shu[z])
					}
				}
			}
			fmt.Println(lspk,lspk1)
			if lspk1[0]==13&&lspk1[1]==4{
				answe.Matches[por].Result=2
			}else if lspk1[0]-lspk1[4]==4{
				answe.Matches[por].Result=2
			}else {
				if lspk[0]==lspk[3]||lspk[1]==lspk[4]{
					answe.Matches[por].Result=1
				}else if lspk[0]==lspk[2]||lspk[2]==lspk[4]{
					if lspk[0]==lspk[2]{
						if lspk[3]==lspk[4]{
							answe.Matches[por].Result=1
						}else {
							answe.Matches[por].Result=2
						}
					}else if lspk[2]==lspk[4]{
						if lspk[0]==lspk[1]{
							answe.Matches[por].Result=1
						}else {
							answe.Matches[por].Result=2
						}
					}
				}else {
					answe.Matches[por].Result=2
				}
			}
		}else {
			fmt.Println(lspk,lspk1)
			for z:=12;z>=0;z--{
				for x:=0;x<10;x=x+2{
					if p[z]==string(v.Alice[x]){
						lspk=append(lspk,shu[z])
					}
					if p[z]==string(v.Bob[x]){
						lspk1=append(lspk1,shu[z])
					}
				}
			}
			fmt.Println(lspk,lspk1)
			lslp=pddj(lspk)
			lslp1=pddj(lspk1)
			if lslp.Ddj!=lslp1.Ddj {
				fmt.Println(lslp.Ddj, lslp1.Ddj)
				fmt.Println(lslp.Pk, lslp1.Pk)
				if lslp.Ddj > lslp1.Ddj {
					answe.Matches[por].Result = 1
				} else {
					answe.Matches[por].Result = 2
				}
			} else{
				fmt.Println(lslp.Ddj,lslp1.Ddj)
				fmt.Println(lslp.Pk,lslp1.Pk)
				switch lslp.Ddj{
				case 7:
					if lslp.Pk[0]!=lslp1.Pk[0]{
						if lslp.Pk[0]>lslp1.Pk[0]{
							answe.Matches[por].Result=1
						} else{
							answe.Matches[por].Result=2
						}
					}else {
						if lslp.Pk[4]!=lslp1.Pk[4]{
							if lslp.Pk[4]>lslp1.Pk[4]{
								answe.Matches[por].Result=1
							} else{
								answe.Matches[por].Result=2
							}
						}else {
							answe.Matches[por].Result=3
						}
					}
				case 6:
					if lslp.Pk[0]!=lslp1.Pk[0]{
						if lslp.Pk[0]>lslp1.Pk[0]{
							answe.Matches[por].Result=1
						} else{
							answe.Matches[por].Result=2
						}
					}else {
						if lslp.Pk[3]!=lslp1.Pk[3]{
							if lslp.Pk[3]>lslp1.Pk[3]{
								answe.Matches[por].Result=1
							} else{
								answe.Matches[por].Result=2
							}
						}else {
							answe.Matches[por].Result=3
						}
					}
				case 5:
					if lspk[0]==13&&lspk[1]==4&&lspk1[0]==13&&lspk1[1]==4{
						answe.Matches[por].Result=3
					}else if lspk[0]==13&&lspk[1]==4{
						answe.Matches[por].Result=2
					}else if lspk1[0]==13&&lspk1[1]==4{
						answe.Matches[por].Result=1
					} else if lspk[0]!=lspk1[0]{
						if lspk[0]>lspk1[0]{
							answe.Matches[por].Result=1
						}else{
							answe.Matches[por].Result=2
						}
					}else{
						answe.Matches[por].Result=3
					}
				case 4:
					if lslp.Pk[0]!=lslp1.Pk[0]{
						if lslp.Pk[0]>lslp1.Pk[0]{
							answe.Matches[por].Result=1
						} else{
							answe.Matches[por].Result=2
						}
					}else {
						if lslp.Pk[3]!=lslp1.Pk[3]{
							if lslp.Pk[3]>lslp1.Pk[3]{
								answe.Matches[por].Result=1
							} else{
								answe.Matches[por].Result=2
							}
						}else {
							if lslp.Pk[4]!=lslp1.Pk[4]{
								if lslp.Pk[4]>lslp1.Pk[4]{
									answe.Matches[por].Result=1
								} else{
									answe.Matches[por].Result=2
								}
							}else{
								answe.Matches[por].Result=3
							}
						}
					}
				case 3:
					if lslp.Pk[0]!=lslp1.Pk[0]{
						if lslp.Pk[0]>lslp1.Pk[0]{
							answe.Matches[por].Result=1
						} else{
							answe.Matches[por].Result=2
						}
					}else {
						if lslp.Pk[2]!=lslp1.Pk[2]{
							if lslp.Pk[2]>lslp1.Pk[2]{
								answe.Matches[por].Result=1
							} else{
								answe.Matches[por].Result=2
							}
						}else {
							if lslp.Pk[4]!=lslp1.Pk[4]{
								if lslp.Pk[4]>lslp1.Pk[4]{
									answe.Matches[por].Result=1
								} else{
									answe.Matches[por].Result=2
								}
							}else{
								answe.Matches[por].Result=3
							}
						}
					}
				case 2:
					if lslp.Pk[0]!=lslp1.Pk[0]{
						if lslp.Pk[0]>lslp1.Pk[0]{
							answe.Matches[por].Result=1
						} else{
							answe.Matches[por].Result=2
						}
					}else {
						if lslp.Pk[2]!=lslp1.Pk[2]{
							if lslp.Pk[2]>lslp1.Pk[2]{
								answe.Matches[por].Result=1
							} else{
								answe.Matches[por].Result=2
							}
						}else {
							if lslp.Pk[3]!=lslp1.Pk[3]{
								if lslp.Pk[3]>lslp1.Pk[3]{
									answe.Matches[por].Result=1
								} else{
									answe.Matches[por].Result=2
								}
							}else{
								if lslp.Pk[4]!=lslp1.Pk[4]{
									if lslp.Pk[4]>lslp1.Pk[4]{
										answe.Matches[por].Result=1
									} else{
										answe.Matches[por].Result=2
									}
								}else {
									answe.Matches[por].Result=3
								}
							}
						}
					}
				case 1:
					if lspk[0]!=lspk1[0]{
						if lspk[0]>lspk1[0]{
							answe.Matches[por].Result=1
						}else{
							answe.Matches[por].Result=2
						}
					}else{
						if lspk[1]!=lspk1[1]{
							if lspk[1]>lspk1[1]{
								answe.Matches[por].Result=1
							}else{
								answe.Matches[por].Result=2
							}
						}else{
							if lspk[2]!=lspk1[2]{
								if lspk[2]>lspk1[2]{
									answe.Matches[por].Result=1
								}else{
									answe.Matches[por].Result=2
								}
							}else{
								if lspk[3]!=lspk1[3]{
									if lspk[3]>lspk1[3]{
										answe.Matches[por].Result=1
									}else{
										answe.Matches[por].Result=2
									}
								}else{
									if lspk[4]!=lspk1[4]{
										if lspk[4]>lspk1[4]{
											answe.Matches[por].Result=1
										}else{
											answe.Matches[por].Result=2
										}
									}else{
										answe.Matches[por].Result=3
									}
								}
							}
						}
					}
				}
			}
		}
		lspk=qk
		lspk1=qk
	}
	t2:=time.Now()
	fmt.Println(t2.Sub(t1))
	filePtr, err := os.Create("answer3.json")
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()
	data, err := json.MarshalIndent(answe, "", "  ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		fmt.Println("Encoder success")
	}

	filePtr.Write(data)

}
