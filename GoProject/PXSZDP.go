package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type APZ struct {
	PZ [5]string
	PZS [5]int
	DX int
	ZL int
}

func main(){
	var i [13]int=[13]int{1,2,3,4,5,6,7,8,9,10,11,12,13}
	var o [13]string = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K","A" }
	//var one [5]string=[5]string{"A","2","3","4","5"}
	//var two [5]string=[5]string{"2","3","4","5","6"}
	//var three [5]string=[5]string{"3","4","5","6","7"}
	//var four [5]string=[5]string{"4","5","6","7","8"}
	//var five [5]string=[5]string{"5","6","7","8","9"}
	//var six [5]string=[5]string{"6","7","8","9","10"}
	//var seven [5]string=[5]string{"7","8","9","10","J"}
	//var eight [5]string=[5]string{"8","9","10","J","Q"}
	//var nine [5]string=[5]string{"9","10","J","Q","k"}
	//var ten [5]string=[5]string{"10","J","Q","K","A"}

	var all [5]string
	var alls [5]int

	var sum int
	var pz []APZ
	var p APZ
	for a:=12;a>=0;a--{
		sum=0
		all[0]=o[a]
		alls[0]=i[a]
		sum=sum+i[a]
		for b:=a-1;b>=0;b--{
			all[1]=o[b]
			alls[1]=i[b]
			sum=sum+i[b]

			for c:=b-1;c>=0;c--{
				all[2]=o[c]
				alls[2]=i[c]
				sum=sum+i[c]
				if c==0 {
					sum=sum-i[c]-i[b]
				}
				for d:=c-1;d>=0;d--{
					all[3]=o[d]
					alls[3]=i[d]
					sum=sum+i[d]
					if d==0 {
						sum=sum-i[d]-i[c]
					}
					for e:=d-1;e>=0;e--{
						all[4]=o[e]
						alls[4]=i[e]
						sum=sum+i[e]

						if all[0]=="A"&&all[1]=="K"&&all[2]=="Q"&&all[3]=="J"&&all[4]=="10"{
							p.PZ=all
							p.PZS=alls
							p.DX=sum
							p.ZL=5
							sum=sum-i[e]
						}else if all[0]=="A"&&all[1]=="5"&&all[2]=="4"&&all[3]=="3"&&all[4]=="2" {
							p.PZ=all
							p.PZS=alls
							p.DX=sum
							p.ZL=5
							sum=sum-i[e]
						}else {
						if all[0]=="6"&&all[4]=="2"||all[0]=="7"&&all[4]=="3"||all[0]=="8"&&all[4]=="4"||all[0]=="9"&&all[4]=="5"||all[0]=="10"&&all[4]=="6"||all[0]=="J"&&all[4]=="7" ||all[0]=="Q"&&all[4]=="8"||all[0]=="K"&&all[4]=="9" {
							p.PZ=all
							p.PZS=alls
							p.DX=sum
							p.ZL=5
							sum=sum-i[e]
						}else {
							p.PZ = all
							p.PZS=alls
							p.DX = sum
							p.ZL = 1
							sum=sum-i[e]
						}}
						if e==0{
							sum=sum-i[d]
						}

						pz=append(pz, p)
					}
				}
			}
		}

	}

	fmt.Println(pz)
	filePtr, err := os.OpenFile("AllCheck4.json",os.O_WRONLY,0644)
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer filePtr.Close()
	data, err := json.MarshalIndent(pz, "", "  ")
	if err != nil {
		fmt.Println("Encoder failed", err.Error())

	} else {
		n,_:=filePtr.Seek(0,os.SEEK_END)
		_,errorp:=filePtr.WriteAt(data,n)
		fmt.Println(errorp)
	}
}









