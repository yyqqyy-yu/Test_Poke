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
var ke [13]string = [13]string{ "2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K","A"}
var ke2 [12]string = [12]string{ "2", "3", "4", "5", "6", "7", "8", "9", "J", "Q", "K","A"}
var shu [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
var ke1 string="23456789TJQKA"
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

func qrth(anyp string,test1 APZ) (res string){

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

func qrth1(anyp string,test1 APZ ) (res string){

		for hs:=3;hs<10;hs=hs+2{
			if string(anyp[1])!=string(anyp[hs]){
				return "no"
				break
			}
		}
		return string(anyp[1])


}

func main() {
	t1 := time.Now()
	//打开文件
	inputFile, inputError := os.Open("AllCheck4.json")
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
	t3:=time.Now()
	fmt.Println(t3.Sub(t1))
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
			for i := 0; i < 5; i++ {
				ae.PZ[i] = string(v.Alice[o])
				o = o + 2
			}
			test1 = append(test1, ae)

			o = 0
			for i := 0; i < 5; i++ {
				bb.PZ[i] = string(v.Bob[o])
				o = o + 2
			}
			test1 = append(test1, bb)
		// fmt.Println("first time to print test1")
			//fmt.Println(test1)
			asum = 0
			for i := 0; i < 5; i++ {
				desum = strings.Index(ke1, test1[0].PZ[i])
				//fmt.Println(desum)
				asum = asum + shu[desum]

			}
			test1[0].DX = asum
			//fmt.Println(asum)
			asum = 0
			for i := 0; i < 5; i++ {
				desum = strings.Index(ke1, test1[1].PZ[i])
				//	fmt.Println(desum)
				asum = asum + shu[desum]
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
						if aflpz[test1[0].DX-6].FPZ[js].PZ[ls] == test1[0].PZ[ls1]{
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
			qrt1=qrth1(answe.Matches[por].Alice,test1[0])
			qrt2=qrth1(answe.Matches[por].Bob,test1[1])
			if qrt1!="no"&&qrt2!="no"||qrt1=="no"&&qrt2=="no"{
				if test1[0].ZL != test1[1].ZL {
					//fmt.Println(test1[1].DX)
					//fmt.Println(test1[1].PZ)
					if test1[0].ZL > test1[1].ZL {
						answe.Matches[por].Result = 1
					} else {
						answe.Matches[por].Result = 2
					}
				} else {
					//fmt.Println(test1[0].ZL)
					switch test1[0].ZL {
					case 7:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 2
							}
						} else {
							if test1[0].PZS[4] != test1[1].PZS[4] {
								if test1[0].PZS[4] > test1[1].PZS[4] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
									answe.Matches[por].Result = 3
							}
						}
					case 6:
						if test1[0].PZS[0]!= test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 2
							}
						} else {
							if test1[0].PZS[3] != test1[1].PZS[3] {
								if test1[0].PZS[3] > test1[1].PZS[3] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
								answe.Matches[por].Result = 3
							}
						}
					case 5:
						huase51=qrth(answe.Matches[por].Alice,test1[0])
						huase52=qrth(answe.Matches[por].Bob,test1[1])
						if test1[0].ZL == 10 && test1[1].ZL == 10 {
							if (test1[0].PZS[0] == 13 && test1[0].PZS[4] == 9) && (test1[1].PZS[0] == 13 && test1[1].PZS[4] == 9) {
								answe.Matches[por].Result = 3
							} else if test1[0].PZS[0] == 13 && test1[0].PZS[4] == 1 {
								answe.Matches[por].Result = 1
							} else if test1[1].PZS[0] == 13 && test1[1].PZS[4] == 1 {
								answe.Matches[por].Result = 2
							} else {
								if test1[0].PZS[0] != test1[1].PZS[0] {
									if test1[0].PZS[0] > test1[1].PZS[0] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 2
									}
								} else {
									answe.Matches[por].Result = 3
								}
							}
						} else if test1[0].ZL == 10 {
							answe.Matches[por].Result = 1
						} else if test1[1].ZL == 10 {
							answe.Matches[por].Result = 2
						} else {
							if test1[0].PZS[0] != test1[1].PZS[0] {
								if test1[0].PZS[0] > test1[1].PZS[0] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
									answe.Matches[por].Result = 3
							}
						}
					case 4:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 2
							}
						} else {
							if test1[0].PZS[3] != test1[1].PZS[3] {
								if test1[0].PZS[3] > test1[1].PZS[3] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
									if test1[0].PZS[4] != test1[1].PZS[4] {
										if test1[0].PZS[4] > test1[1].PZS[4] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 2
										}
									}else{
										answe.Matches[por].Result = 3
									}
							}
						}
					case 3:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 2
							}
						} else {
							if test1[0].PZS[2] != test1[1].PZS[2] {
								if test1[0].PZS[2] > test1[1].PZS[2] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
								if test1[0].PZ[4] != test1[1].PZ[4] {
									if test1[0].PZ[4] > test1[1].PZ[4] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 2
									}
								} else {
										answe.Matches[por].Result = 3
								}
							}
						}
					case 2:
						if test1[0].PZS[0] != test1[1].PZS[0] {
							if test1[0].PZS[0] > test1[1].PZS[0] {
								answe.Matches[por].Result = 1
							} else {
								answe.Matches[por].Result = 2
							}
						} else {
							if test1[0].PZS[2] != test1[1].PZS[2] {
								if test1[0].PZS[2] > test1[1].PZS[2] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
								if test1[0].PZS[3] != test1[1].PZS[3] {
									if test1[0].PZS[3] > test1[1].PZS[3] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 2
									}
								} else {
									if test1[0].PZS[4] != test1[1].PZS[4] {
										if test1[0].PZS[4] > test1[1].PZS[4] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 2
										}
									} else {
										answe.Matches[por].Result = 3
									}
								}
							}
						}
					case 1:
						//fmt.Println(test1)
						huase11 = qrth(answe.Matches[por].Alice,  test1[0])
						huase12 = qrth(answe.Matches[por].Bob, test1[1])
						if test1[0].ZL == 10 && test1[1].ZL == 10 {
							if test1[0].PZS[0] != test1[1].PZS[0] {
								if test1[0].PZS[0] > test1[1].PZS[0] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							}else {
								if test1[0].PZS[1] != test1[1].PZS[1] {
									if test1[0].PZS[1] > test1[1].PZS[1] {
										answe.Matches[por].Result = 1
									} else {
										answe.Matches[por].Result = 2
									}
								}else {
									if test1[0].PZS[2] != test1[1].PZS[2] {
										if test1[0].PZS[2] > test1[1].PZS[2] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 2
										}
									}else{
										if test1[0].PZS[3] != test1[1].PZS[3] {
											if test1[0].PZS[3] > test1[1].PZS[3] {
												answe.Matches[por].Result = 1
											} else {
												answe.Matches[por].Result = 2
											}
										}else{
											if test1[0].PZS[4] != test1[1].PZS[4] {
												if test1[0].PZS[4] > test1[1].PZS[4] {
													answe.Matches[por].Result = 1
												} else {
													answe.Matches[por].Result = 2
												}
											}else{
												answe.Matches[por].Result = 3
											}
										}
									}
								}
							}
						} else if test1[0].ZL == 10 {
							answe.Matches[por].Result = 1
						} else if test1[1].ZL == 10 {
							answe.Matches[por].Result = 2
						} else {
							//fmt.Println(test1)
							if test1[0].PZS[0] != test1[1].PZS[0] {
								if test1[0].PZS[0] > test1[1].PZS[0] {
									answe.Matches[por].Result = 1
								} else {
									answe.Matches[por].Result = 2
								}
							} else {
									if test1[0].PZS[1] != test1[1].PZS[1] {
										if test1[0].PZS[1] > test1[1].PZS[1] {
											answe.Matches[por].Result = 1
										} else {
											answe.Matches[por].Result = 2
										}
									} else {
											if test1[0].PZS[2] != test1[1].PZS[2] {
												if test1[0].PZS[2] > test1[1].PZS[2] {
													answe.Matches[por].Result = 1
												} else {
													answe.Matches[por].Result = 2
												}
											} else {
													if test1[0].PZS[3] != test1[1].PZS[3] {
														if test1[0].PZS[3] > test1[1].PZS[3] {
															answe.Matches[por].Result = 1
														} else {
															answe.Matches[por].Result = 2
														}
													} else {
															if test1[0].PZS[4] != test1[1].PZS[4] {
																if test1[0].PZS[4] > test1[1].PZS[4] {
																	answe.Matches[por].Result = 1
																} else {
																	answe.Matches[por].Result = 2
																}
															} else {
																	answe.Matches[por].Result = 3
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
						answe.Matches[por].Result = 2
					} else {
						answe.Matches[por].Result = 1
					}
			}else if qrt1=="no" &&qrt2!="no"{
				if test1[0].DX ==7||test1[0].DX ==8{
					answe.Matches[por].Result = 1
				} else {
					answe.Matches[por].Result = 2
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
			filePtr, err := os.Create("answer1.json")
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