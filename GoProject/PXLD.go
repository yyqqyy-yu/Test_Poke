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
	var i [13]int = [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	var o [13]string = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K","A" }

	var all [5]string
	var alls [5]int

	var sum int
	var pz []APZ
	var p APZ
	for a := 12; a >=0; a-- {
		sum = 0
		all[0] = o[a]
		alls[0]=i[a]
		all[1] = o[a]
		alls[1]=i[a]


		sum = sum + 2*i[a]
		for b := a-1; b >=0; b-- {
			all[2]=o[b]
			alls[2]=i[b]
			all[3]=o[b]
			alls[3]=i[b]
			sum=sum+2*i[b]

			for c:=0;c<len(o);c++{
				if(c!=a&&c!=b){
					sum=sum+i[c]
					all[4]=o[c]
					alls[4]=i[c]
					p.PZ=all
					p.PZS=alls
					p.DX=sum
					p.ZL=3
					sum=sum-i[c]
					pz=append(pz, p)

				}
				if c==12{
					sum=sum-2*i[b]
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
