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
	for a := 0; a < len(o); a++ {
		sum = 0
		all[0] = o[a]
		alls[0]=i[a]
		all[1] = o[a]
		alls[1]=i[a]
		all[2] = o[a]
		alls[2]=i[a]

		sum = sum + 3*i[a]
		for b := 12; b >=0; b-- {
			if b!=a {
				all[3]=o[b]
				alls[3]=i[b]

			sum=sum+i[b]

			for c:=b-1;c>=0;c--{
				if c!=a {
					all[4]=o[c]
					alls[4]=i[c]

				sum=sum+i[c]
				p.PZ=all
				p.PZS=alls
				p.DX=sum
				p.ZL=4
				pz=append(pz, p)
				sum=sum-i[c]
				}
				if c==0{
					sum=sum-i[b]
				}
			}
			}
			if b==0{
				sum=sum-i[b]
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




