package main

import (
	"bufio"
	"time"

	"encoding/json"

	"strings"

	//"strings"

	"fmt"
	"io"
	"os"
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
var qrt1 string
var qrt2 string
var ke [13]string = [13]string{ "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K","A"}
var ke2 [12]string = [12]string{ "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K","A"}
var shu [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var ke1 string="2345678910JQKA"
var desum int
var asum int
var data []byte
var err error
var answe APP
var jilu []string
var jilu1 []string
var qkjl []string
var test1 []APZ
var testt []APZ
var ae APZ
var bb APZ
var o int
var f int
var d int
var t int
var s string
var js string
var e string
var all string
var calcu int
var huase string
var huas int
var huase1 string
var huas1 int
var huase51 string
var huas51 int
var huase52 string
var huas52 int
var huase41 string
var huas41 int
var huase42 string
var huas42 int
var huase43 string
var huas43 int
var huase44 string
var huas44 int
var huase31 string
var huas31 int
var huase32 string
var huas32 int
var huase21 string
var huas21 int
var huase22 string
var huas22 int
var huase23 string
var huas23 int
var huase24 string
var huas24 int
var huase25 string
var huas25 int
var huase26 string
var huas26 int
var huase11 string
var huas11 int
var huase12 string
var huas12 int
var huase13 string
var huas13 int
var huase14 string
var huas14 int
var huase15 string
var huas15 int
var huase16 string
var huas16 int
var huase17 string
var huas17 int
var huase18 string
var huas18 int
var huase19 string
var huas19 int
var huase120 string
var huas120 int
var huase121 string
var huas121 int
var huase122 string
var huas122 int
var tenpz string
var tenpz1 string


func pdhs(huase string) (check int){
	if huase=="s"{
		return 4
	} else if huase=="h"{
		return 3
	} else if huase=="d"{
		return 2
	} else if huase=="c"{
		return 1
	}
	return 0
}

func findhs(alices string,char string) (res string){
	return string(alices[strings.IndexAny(alices,char)+1])
}

func qrth(anyp string,tenpz string,test1 APZ, ) (res string){
	if len(anyp)>10{
		for hs:=3;hs<10;hs=hs+2{
			if string(tenpz[1])!=string(tenpz[hs]){
				test1.ZL=1
				break
			}else {
				test1.ZL=10
			}
		}
		return string(tenpz[1])
	}else {
		for hs:=3;hs<10;hs=hs+2{
			if string(anyp[1])!=string(anyp[hs]){
				test1.ZL=5
				break
			}else {
				test1.ZL=10
			}
		}
		return string(anyp[1])
	}

}

func qrth1(anyp string,tenpz string,test1 APZ ) (res string){
	if len(anyp)>10{
		for hs:=3;hs<10;hs=hs+2{
			if string(tenpz[1])!=string(tenpz[hs]){
				return "no"
				break
			}
		}
		return string(tenpz[1])
	}else {
		for hs:=3;hs<10;hs=hs+2{
			if string(anyp[1])!=string(anyp[hs]){
				return "no"
				break
			}
		}
		return string(anyp[1])
	}

}

func main() {
	t1 := time.Now()
	//打开文件
	inputFile, inputError := os.Open("AllCheck3.json")
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
	var aflpz []FLPZ
	var flpz FLPZ
	var test FLPZ
	//po:=
		json.Unmarshal([]byte(js), &ans)
	//fmt.Println(po)
	//fmt.Println(ans)
	for i := 6; i <= 64; i++ {
		flpz = test
		for _, v := range ans {
			if v.DX == i {
				flpz.FPZ = append(flpz.FPZ, v)
			}
		}

		aflpz = append(aflpz, flpz)

	}
//fmt.Println(aflpz)

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
	fmt.Println("fyhrtyrtyrtyryrtyrtyrty")
		//fmt.Println(answe)
		for por, v := range answe.Matches {
			o = 0
			if strings.Contains(v.Alice, "0") {
				if len(v.Alice) == 11 {
					f = strings.Index(v.Alice, "0")
					s = v.Alice[0:f]
					e = v.Alice[f+1:]
					all = s + e
					tenpz = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							ae.PZ[i] = "10"
							o = o + 2
						} else {
							ae.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}

				} else if len(v.Alice) == 12 {
					f = strings.Index(v.Alice, "0")
					s = v.Alice[0:f]
					e = v.Alice[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					tenpz = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							ae.PZ[i] = "10"
							o = o + 2
						} else {
							ae.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}

				} else if len(v.Alice) == 13 {
					f = strings.Index(v.Alice, "0")
					s = v.Alice[0:f]
					e = v.Alice[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					tenpz = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							ae.PZ[i] = "10"
							o = o + 2
						} else {
							ae.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}
				} else if len(v.Alice) > 14 {
var bingo int=0
					for st := 0; st < len(ke2); st++ {

						bingo := strings.IndexAny(v.Alice, ke2[st])
						if bingo > 0 {
							for i := 0; i < 4; i++ {
								ae.PZ[i] = "10"
							}
							ae.PZ[4] = string(v.Alice[bingo])
						}
					}
					if bingo==12{
						for lss:=2;lss<12;lss=lss+3{
							tenpz=tenpz+"1"+string(v.Alice[lss])
						}

						tenpz=tenpz+v.Alice[12:]
					}else if bingo==0{
						for lss:=4;lss<14;lss=lss+3{
							tenpz=tenpz+"1"+string(v.Alice[lss])
						}
						tenpz=tenpz+v.Alice[:2]
					}else {
						s = all[0:bingo]
						e = all[f+2:]
						all=s+e
						for lss:=4;lss<14;lss=lss+3{
							tenpz=tenpz+"1"+string(all[lss])
						}
						tenpz=tenpz+v.Alice[bingo:bingo+2]
					}
				}
			} else {
				for i := 0; i < 5; i++ {
					ae.PZ[i] = (string(v.Alice[o]))
					o = o + 2
				}
			}
			test1 = append(test1, ae)
			o = 0
			if strings.Contains(v.Bob, "0") {
				if len(v.Bob) == 11 {
					f = strings.Index(v.Bob, "0")
					s = v.Bob[0:f]
					e = v.Bob[f+1:]
					all = s + e
					tenpz1 = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							bb.PZ[i] = "10"
							o = o + 2
						} else {
							bb.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}
				} else if len(v.Bob) == 12 {
					f = strings.Index(v.Bob, "0")
					s = v.Bob[0:f]
					e = v.Bob[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					tenpz1 = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							bb.PZ[i] = "10"
							o = o + 2
						} else {
							bb.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}
				} else if len(v.Bob) == 13 {
					f = strings.Index(v.Bob, "0")
					s = v.Bob[0:f]
					e = v.Bob[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					f = strings.Index(all, "0")
					s = all[0:f]
					e = all[f+1:]
					all = s + e
					tenpz1 = all
					for i := 0; i < 5; i++ {
						if string(all[o]) == "1" {
							bb.PZ[i] = "10"
							o = o + 2
						} else {
							bb.PZ[i] = (string(all[o]))
							o = o + 2
						}
					}
				} else if len(v.Bob) == 14 {
					var bingo int=0
					for st := 0; st < len(ke2); st++ {
						bingo := strings.IndexAny(v.Bob, ke2[st])
						if bingo > 0 {
							for i := 0; i < 4; i++ {
								bb.PZ[i] = "10"
							}
							bb.PZ[4] = string(v.Bob[bingo])
						}
					}
					if bingo==12{
						for lss:=2;lss<12;lss=lss+3{
							tenpz1=tenpz1+"1"+string(v.Bob[lss])
						}

						tenpz1=tenpz1+v.Bob[12:]
					}else if bingo==0{
						for lss:=4;lss<14;lss=lss+3{
							tenpz1=tenpz1+"1"+string(v.Bob[lss])
						}
						tenpz1=tenpz1+v.Bob[:2]
					}else {
						s = all[0:bingo]
						e = all[f+2:]
						all=s+e
						for lss:=4;lss<14;lss=lss+3{
							tenpz1=tenpz1+"1"+string(all[lss])
						}
						tenpz1=tenpz1+v.Bob[bingo:bingo+2]
					}


				}
			} else {
				for i := 0; i < 5; i++ {
					bb.PZ[i] = string(v.Bob[o])
					o = o + 2
				}
			}
			test1 = append(test1, bb)
		// fmt.Println("first time to print test1")
			//fmt.Println(test1)
			asum = 0
			for i := 0; i < 5; i++ {
				desum = strings.Index(ke1, test1[0].PZ[i])
				//fmt.Println(desum)
				if test1[0].PZ[i] == "A"|| test1[0].PZ[i] == "Q"|| test1[0].PZ[i] == "J"||test1[0].PZ[i] == "K"{
					asum = asum + shu[desum-1]
				} else {
					asum = asum + shu[desum]
				}
			}
			test1[0].DX = asum
			//fmt.Println(asum)
			asum = 0
			for i := 0; i < 5; i++ {
				desum = strings.Index(ke1, test1[1].PZ[i])
				//	fmt.Println(desum)
				if test1[1].PZ[i] == "A"|| test1[1].PZ[i] == "K"|| test1[1].PZ[i] == "Q"||test1[1].PZ[i] == "J"{
					asum = asum + shu[desum-1]
				} else {
					//fmt.Println(test1[1].PZ[i])
					asum = asum + shu[desum]
				}
			}
			test1[1].DX = asum
			//fmt.Println(asum)
			//fmt.Println(test1)
			for jl := 0; jl < 5; jl++ {
				jilu = append(jilu, test1[0].PZ[jl])
				jilu1 = append(jilu1, test1[1].PZ[jl])
			}
			calcu = 0
			for js := 0; js < len(aflpz[test1[0].DX-6].FPZ); js++ {
				calcu = 0
				for ls := 0; ls < 5; ls++ {
					for ls1 := 0; ls1 < 5; ls1++ {
						if aflpz[test1[0].DX-6].FPZ[js].PZ[ls] == test1[0].PZ[ls1] {
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
						test1[0] = aflpz[test1[0].DX-6].FPZ[js]
						//fmt.Println("chu xian =0")
						break
					}
				}
				calcu = 0
				for js := 0; js < len(aflpz[test1[1].DX-6].FPZ); js++ {
					calcu = 0
					for ls := 0; ls < 5; ls++ {
						for ls1 := 0; ls1 < 5; ls1++ {
							if aflpz[test1[1].DX-6].FPZ[js].PZ[ls] == test1[1].PZ[ls1] {
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
						test1[1] = aflpz[test1[1].DX-6].FPZ[js]
						//fmt.Println("chu xian =1")
						break
					}
				}
				//fmt.Println(tenpz)
			qrt1=qrth1(answe.Matches[por].Alice,tenpz,test1[0])
			qrt2=qrth1(answe.Matches[por].Bob,tenpz1,test1[1])
			if qrt1!="no"&&qrt2!="no"||qrt1=="no"&&qrt2=="no"{
				if test1[0].ZL != test1[1].ZL {
					//fmt.Println(test1[1].DX)
					//fmt.Println(test1[1].PZ)
					if test1[0].ZL > test1[1].ZL {
						answe.Matches[por].Result = 1
					} else {
						answe.Matches[por].Result = 0
					}
				} else {
					//fmt.Println(test1[0].ZL)
					switch test1[0].ZL {
					case 7:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 0
							}
						} else {
							if test1[0].PZS[4] != test1[1].PZS[4] {
								if test1[0].PZS[4] > test1[1].PZS[4] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								huase = string(answe.Matches[por].Alice[strings.IndexAny(answe.Matches[por].Alice, test1[0].PZ[4])+1])
								huase1 = string(answe.Matches[por].Bob[strings.IndexAny(answe.Matches[por].Bob, test1[1].PZ[4])+1])
								if huase == huase1 {
									answe.Matches[por].Result = 2
								} else {
									huas = pdhs(huase)
									huas1 = pdhs(huase1)
									if huas > huas1 {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								}
							}
						}
					case 6:
						//.Println(test1[0].PZS[0])
						if test1[0].PZS[0]!= test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 0
							}
						} else {
							if test1[0].PZS[3] != test1[1].PZS[3] {
								if test1[0].PZS[3] > test1[1].PZS[3] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								answe.Matches[por].Result = 2
							}
						}
					case 5:
						if len(answe.Matches[por].Alice) > 10 {
							for hs := 3; hs < 10; hs = hs + 2 {
								if string(tenpz[1]) != string(tenpz[hs]) {
									test1[0].ZL = 5
									break
								} else {
									test1[0].ZL = 10
								}
							}
							huase51 = string(tenpz[1])
						} else {
							for hs := 3; hs < 10; hs = hs + 2 {
								if string(answe.Matches[por].Alice[1]) != string(answe.Matches[por].Alice[hs]) {
									test1[0].ZL = 5
									break
								} else {
									test1[0].ZL = 10
								}
							}
							huase51 = string(answe.Matches[por].Alice[1])
						}
						if len(answe.Matches[por].Bob) > 10 {
							for hs := 3; hs < 10; hs = hs + 2 {
								if string(tenpz1[1]) != string(tenpz1[hs]) {
									test1[1].ZL = 5
									break
								} else {
									test1[1].ZL = 10
								}
							}
							huase52 = string(tenpz1[1])
						} else {
							for hs := 3; hs < 10; hs = hs + 2 {
								if string(answe.Matches[por].Bob[1]) != string(answe.Matches[por].Bob[hs]) {
									test1[1].ZL = 5
									break
								} else {
									test1[1].ZL = 10
								}
							}
							huase52 = string(answe.Matches[por].Bob[1])
						}
						if test1[0].ZL == 10 && test1[1].ZL == 10 {
							if (test1[0].PZS[0] == 13 && test1[0].PZS[4] == 1) && (test1[1].PZS[0] == 13 && test1[1].PZS[4] == 1) {
								answe.Matches[por].Result = 2
							} else if test1[0].PZS[0] == 13 && test1[0].PZS[4] == 1 {
								answe.Matches[por].Result = 1
							} else if test1[1].PZS[0] == 13 && test1[1].PZS[4] == 1 {
								answe.Matches[por].Result = 0
							} else {
								if test1[0].PZS[0] != test1[1].PZS[0] {
									if test1[0].PZS[0] > test1[1].PZS[0] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									if huase51 != huase52 {
										huas51 = pdhs(huase51)
										huas52 = pdhs(huase52)
										if huas51 > huas52 {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										answe.Matches[por].Result = 2
									}
								}
							}

						} else if test1[0].ZL == 10 {
							answe.Matches[por].Result = 1
						} else if test1[1].ZL == 10 {
							answe.Matches[por].Result = 0
						} else {
							if test1[0].PZS[0] != test1[1].PZS[0] {
								if test1[0].PZS[0] > test1[1].PZS[0] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								if huase51 != huase52 {
									huas51 = pdhs(huase51)
									huas52 = pdhs(huase52)
									if huas51 > huas52 {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									answe.Matches[por].Result = 2
								}

							}
						}
					case 4:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 0
							}
						} else {
							if test1[0].PZS[3] != test1[1].PZS[3] {
								if test1[0].PZS[3] > test1[1].PZS[3] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								huase41 = findhs(answe.Matches[por].Alice, test1[0].PZ[3])
								huase42 = findhs(answe.Matches[por].Bob, test1[1].PZ[3])
								huas41 = pdhs(huase41)
								huas42 = pdhs(huase42)
								if huas41 != huas42 {
									if huas41 > huas42 {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									if test1[0].PZS[4] != test1[1].PZS[4] {
										if test1[0].PZS[4] > test1[1].PZS[4] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										huase43 = findhs(answe.Matches[por].Alice, test1[0].PZ[4])
										huase44 = findhs(answe.Matches[por].Bob, test1[1].PZ[4])
										huas43 = pdhs(huase43)
										huas44 = pdhs(huase44)
										if huas43 != huas44 {
											if huas43 > huas44 {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 0
											}
										} else {
											answe.Matches[por].Result = 2
										}
									}
								}
							}
						}
					case 3:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 0
							}
						} else {
							if test1[0].PZS[2] != test1[1].PZS[2] {
								if test1[0].PZS[2] > test1[1].PZS[2] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								if test1[0].PZ[4] != test1[1].PZ[4] {
									if test1[0].PZ[4] > test1[1].PZ[4] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									huase31 = findhs(answe.Matches[por].Alice, test1[0].PZ[4])
									huase32 = findhs(answe.Matches[por].Bob, test1[1].PZ[4])
									huas31 = pdhs(huase31)
									huas32 = pdhs(huase32)
									if huas31 != huas32 {
										if huas31 > huas32 {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										answe.Matches[por].Result = 2
									}
								}

							}
						}
					case 2:
						//fmt.Println(test1)
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 0
							}
						} else {
							if test1[0].PZS[2] != test1[1].PZS[2] {
								if test1[0].PZS[2] > test1[1].PZS[2] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								huase21 = findhs(answe.Matches[por].Alice, test1[0].PZ[2])
								huase22 = findhs(answe.Matches[por].Bob, test1[1].PZ[2])
								huas21 = pdhs(huase21)
								huas22 = pdhs(huase22)
								if huas21 != huas22 {
									if huas21 > huas22 {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									if test1[0].PZS[3] != test1[1].PZS[3] {
										if test1[0].PZS[3] > test1[1].PZS[3] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										huase23 = findhs(answe.Matches[por].Alice, test1[0].PZ[3])
										huase24 = findhs(answe.Matches[por].Bob, test1[1].PZ[3])
										huas23 = pdhs(huase23)
										huas24 = pdhs(huase24)
										if huas23 != huas24 {
											if huas23 > huas24 {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 0
											}
										} else {
											if test1[0].PZS[4] != test1[1].PZS[4] {
												if test1[0].PZS[4] > test1[1].PZS[4] {
													answe.Matches[por].Result = 1
												} else {
													answe.Matches[por].Result = 0
												}
											} else {
												huase25 = findhs(answe.Matches[por].Alice, test1[0].PZ[4])
												huase26 = findhs(answe.Matches[por].Bob, test1[1].PZ[4])
												huas25 = pdhs(huase25)
												huas26 = pdhs(huase26)
												if huas25 != huas26 {
													if huas25 > huas26 {
														answe.Matches[por].Result = 1
													} else {
														answe.Matches[por].Result = 0
													}
												} else {
													answe.Matches[por].Result = 2
												}
											}
										}
									}
								}
							}
						}
					case 1:
						//fmt.Println(test1)
						huase11 = qrth(answe.Matches[por].Alice, tenpz, test1[0])
						huase12 = qrth(answe.Matches[por].Bob, tenpz1, test1[1])
						huas11 = pdhs(huase11)
						huas12 = pdhs(huase12)

						if test1[0].ZL == 10 && test1[1].ZL == 10 {
							if huas11 != huas12 {
								if huas11 > huas12 {
									if test1[0].PZS[0] != test1[1].PZS[0] {
										if test1[0].PZS[0] > test1[1].PZS[0] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										answe.Matches[por].Result = 1
									}
								} else {
									if huas11 < huas12 {
										if test1[0].PZS[0] != test1[1].PZS[0] {
											if test1[0].PZS[0] > test1[1].PZS[0] {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 0
											}
										} else {
											answe.Matches[por].Result = 0
										}
									}
								}
							} else {
								if test1[0].PZS[0] != test1[1].PZS[0] {
									if test1[0].PZS[0] > test1[1].PZS[0] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									if test1[0].PZS[1] != test1[1].PZS[1] {
										if test1[0].PZS[1] > test1[1].PZS[1] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										if test1[0].PZS[2] != test1[1].PZS[2] {
											if test1[0].PZS[2] > test1[1].PZS[2] {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 0
											}
										} else {
											if test1[0].PZS[3] != test1[1].PZS[3] {
												if test1[0].PZS[3] > test1[1].PZS[3] {
													answe.Matches[por].Result = 1
												} else {
													answe.Matches[por].Result = 0
												}
											} else {
												if test1[0].PZS[4] != test1[1].PZS[4] {
													if test1[0].PZS[4] > test1[1].PZS[4] {
														answe.Matches[por].Result = 1
													} else {
														answe.Matches[por].Result = 0
													}
												}
											}
										}
									}
								}
							}
						} else if test1[0].ZL == 10 {
							answe.Matches[por].Result = 1
						} else if test1[1].ZL == 10 {
							answe.Matches[por].Result = 0
						} else {
							fmt.Println(test1)
							if test1[0].PZS[0] != test1[1].PZS[0] {
								if test1[0].PZS[0] > test1[1].PZS[0] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 0
								}
							} else {
								huase13 = findhs(answe.Matches[por].Alice, test1[0].PZ[0])
								huase14 = findhs(answe.Matches[por].Bob, test1[1].PZ[0])
								huas13 = pdhs(huase13)
								huas14 = pdhs(huase14)
								if huas13 != huas14 {
									if huas13 > huas14 {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 0
									}
								} else {
									if test1[0].PZS[1] != test1[1].PZS[1] {
										if test1[0].PZS[1] > test1[1].PZS[1] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 0
										}
									} else {
										huase15 = findhs(answe.Matches[por].Alice, test1[0].PZ[1])
										huase16 = findhs(answe.Matches[por].Bob, test1[1].PZ[1])
										huas15 = pdhs(huase13)
										huas16 = pdhs(huase14)
										if huas15 != huas16 {
											if huas15 > huas16 {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 0
											}
										} else {
											if test1[0].PZS[2] != test1[1].PZS[2] {
												if test1[0].PZS[2] > test1[1].PZS[2] {
													answe.Matches[por].Result = 1
												} else {
													answe.Matches[por].Result = 0
												}
											} else {
												huase17 = findhs(answe.Matches[por].Alice, test1[0].PZ[2])
												huase18 = findhs(answe.Matches[por].Bob, test1[1].PZ[2])
												huas17 = pdhs(huase17)
												huas18 = pdhs(huase18)
												if huas17 != huas18 {
													if huas17 > huas18 {
														answe.Matches[por].Result = 1
													} else {
														answe.Matches[por].Result = 0
													}
												} else {
													if test1[0].PZS[3] != test1[1].PZS[3] {
														if test1[0].PZS[3] > test1[1].PZS[3] {
															answe.Matches[por].Result = 1
														} else {
															answe.Matches[por].Result = 0
														}
													} else {
														huase19 = findhs(answe.Matches[por].Alice, test1[0].PZ[3])
														huase120 = findhs(answe.Matches[por].Bob, test1[1].PZ[3])
														huas19 = pdhs(huase19)
														huas120 = pdhs(huase120)
														if huas19 != huas120 {
															if huas19 > huas120 {
																answe.Matches[por].Result = 1
															} else {
																answe.Matches[por].Result = 0
															}
														} else {
															if test1[0].PZS[4] != test1[1].PZS[4] {
																if test1[0].PZS[4] > test1[1].PZS[4] {
																	answe.Matches[por].Result = 1
																} else {
																	answe.Matches[por].Result = 0
																}
															} else {
																huase121 = findhs(answe.Matches[por].Alice, test1[0].PZ[4])
																huase122 = findhs(answe.Matches[por].Bob, test1[1].PZ[4])
																huas121 = pdhs(huase121)
																huas122 = pdhs(huase122)
																if huas21 != huas122 {
																	if huas121 > huas122 {
																		answe.Matches[por].Result = 1
																	} else {
																		answe.Matches[por].Result = 0
																	}
																} else {
																	answe.Matches[por].Result = 2
																}
															}
														}

													}
												}

											}

										}

									}
								}
							}
						}
					default:
						fmt.Println("moren~~~~~~~~~~~~~~~~~~~~~")
						fmt.Println(test1)
					}
				}
			}else if qrt1!="no"&&qrt2=="no" {
				if test1[1].DX ==7||test1[1].DX ==8{
						answe.Matches[por].Result = 0
					} else {
						answe.Matches[por].Result = 1
					}
			}else if qrt1=="no" &&qrt2!="no"{
				if test1[0].DX ==7||test1[0].DX ==8{
					answe.Matches[por].Result = 1
				} else {
					answe.Matches[por].Result = 0
				}
			}



				//fmt.Println(test1)

				test1 = testt
				jilu = qkjl
				jilu1 = qkjl
			}
	t2 := time.Now()


	fmt.Println(t2.Sub(t1))
			//fmt.Println(answe)
			filePtr, err := os.Create("answer.json")
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