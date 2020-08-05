package main

import (

	"fmt"
)

func main(){
	var k byte
	var j byte=67
	var e string
	var i [13]byte=[13]byte{'A','2','3','4','5','6','7','8','9','0','J','Q','K'}
	var o [13]string=[13]string{"A","2","3","4","5","6","7","8","9","10","J","Q","K"}
	//var o byte='2'
	fmt.Println(k)
	for _,v:=range i{
		k=k+v

	fmt.Println(k)}
if k>j {
	fmt.Println("fuck you")
}
	for _,v:=range o{
	e=e+v
		fmt.Println(e)
		fmt.Println(v)}
}
