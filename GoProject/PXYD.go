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

	var sum int=0
	var pz []APZ
	var p APZ
	for a := 0; a <len(o); a++ {
		sum = 0
		all[0] = o[a]
		all[1] = o[a]
		alls[0]=i[a]
		alls[1]=i[a]
		sum = sum + 2*i[a]
		for b := 12; b>=0; b--{
			if b!=a{
				all[2]=o[b]
				alls[2]=i[b]
				sum=sum+i[b]
				for c:=b-1;c>=0;c-- {
					if c!=a{
						all[3]=o[c]
						alls[3]=i[c]
						if c!=0{
						sum=sum+i[c]}
						for d:=c-1;d>=0;d--{
							if d!=a{
								all[4]=o[d]
								alls[4]=i[d]
								sum=sum+i[d]
								if c==11 {fmt.Println(i[d])}
								p.PZ=all
								p.PZS=alls
								p.DX=sum
								p.ZL=2
								sum=sum-i[d]
								pz=append(pz, p)
							}
								if d==0{
									sum=sum-i[c]
								}
						}
					}
					if c==0{
						sum=sum-i[b]

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

	//filePtr.Write(data)

}
