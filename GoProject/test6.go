package main

import "fmt"

var one [5]string=[5]string{"A","2","3","4","5"}
var two [5]string=[5]string{"A","3","2","4","5"}

func main(){

	if one==two{
		fmt.Println("this is right")
	}

}
