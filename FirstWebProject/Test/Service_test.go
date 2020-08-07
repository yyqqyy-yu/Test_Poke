package gotest

import (
	"db/createTable.go/model"
	"db/createTable.go/service"
	"fmt"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	orList:=service.SelA()
	if orList!=nil{
		fmt.Println(orList)
	}else{
		fmt.Println(orList)
		t.Errorf("wrong")
	}
}

func Test_2(t *testing.T) {
	orList,err:=service.SelByMY()
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println("it's wrong")
	}else{

		fmt.Println(orList)
	}
}
func Test_3(t *testing.T) {
	orList,err:=service.SelByCT()
	if err!=nil{
		fmt.Println(err)
		t.Errorf("wrong")
	}else{
		fmt.Println(orList)
	}
}
func Test_4(t *testing.T) {
	err,or:=service.SelOne(1)
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println(err)
	}else{
		fmt.Println(or)
	}
}
func Test_5(t *testing.T) {
	var or model.Order
	or.ID=25
	err:=service.InsO(or)
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println(err)
	}else{
		fmt.Println("success to add")
	}
}
func Test_6(t *testing.T) {
	err:=service.DelO(25)
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println(err)
	}else{
		fmt.Println("success to del")
	}
}

func Test_7(t *testing.T) {
	var or model.Order
	or.File_url="call me liangzai"
	or.ID=1
	err:=service.UpdO(or)
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println(err)
	}else{
		fmt.Println("success to update")
	}
}

func Test_8(t *testing.T) {
	var or model.Order
	or.ID=uint(time.Now().Unix())

	err:=service.InsN(or)
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println(err)
	}else{
		fmt.Println("success to rollback")
	}
}

func Test_9(t *testing.T) {


	orList:=service.SelLike("5")
	if orList!=nil{
		fmt.Println(orList)
	}else{
		t.Errorf("wrong")
		fmt.Println("fail to sel")
	}
}

func Test_10(t *testing.T) {


	err:=service.SaveExcel
	if err!=nil{
		t.Errorf("wrong")
		fmt.Println("fail to create execl")
	}else{
		fmt.Println("success to create execl")
	}
}

