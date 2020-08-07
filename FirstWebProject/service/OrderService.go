package service

import (
	dbb "db/createTable.go/db"
	"db/createTable.go/model"
	"fmt"
	"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
	"strconv"
)

var db *gorm.DB=dbb.Minit("/home/yyq/FirstWebProject/db/db.yaml")

func UpdO(or model.Order)(err error){

	return dbb.Upd(dbb.DB,or)
}
func InsO(or model.Order)(err error){
	return dbb.CreateB(dbb.DB,or)
}
func DelO(id uint)(err error){
	return dbb.Del(dbb.DB,id)
}
func SelA()(orList []model.Order){
	return dbb.SelA(dbb.DB)
}

func SelO(id uint,or model.Order) (err error){
	return dbb.SelO(dbb.DB,id,or)
}

func SelOne(id uint) (err error,bor  model.Order){
	return dbb.SelOne(dbb.DB,id)
}

func SelByCT() (orList []model.Order,err error){
	return dbb.SelByCT(dbb.DB)
}
func SelByMY() (orList []model.Order,err error){
	return dbb.SelByMY(dbb.DB)
}

func InsN(or model.Order ) (err error){
	tx:= dbb.DB.Begin()

	err=dbb.CreateB(tx,or)
	or.ID=1
	err=dbb.CreateB(tx,or)
	//result:=tx.Exec("INSERT INTO orders (id,user_name,file_url) VALUES (?, ?,?)",25,"33333","it")
	//fmt.Println(result)
	// result=tx.Exec("INSERT INTO orders (id,user_name,file_url) VALUES (?, ?,?)",25,"33333","it")
	if err!=nil{
		tx.Rollback()
	}
	//or.ID=20
	// dbb.CreateB(tx,or)
	tx.Commit()
	return err
}
func SelLike (user_name string)(orList []model.Order){

	orList=dbb.SelLike(dbb.DB,user_name)
	return orList
}

func SaveExcel()(err error){
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row  *xlsx.Row
	var row1  *xlsx.Row
	var rows []*xlsx.Row
	var cell *xlsx.Cell
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		//return err
	}
	orList:=SelA()
	row1=sheet.AddRow()
	for i,_:=range orList{
		fmt.Println(i)
		row=sheet.AddRow()
		rows=append(rows, row)
	}
	row=row1
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "CreatedAt"
	cell = row.AddCell()
	cell.Value = "UpdatedAt"
	cell = row.AddCell()
	cell.Value = "DeletedAt"
	cell = row.AddCell()
	cell.Value = "Order_no"
	cell = row.AddCell()
	cell.Value = "User_name"
	cell = row.AddCell()
	cell.Value = "Amount"
	cell = row.AddCell()
	cell.Value = "Status"
	cell = row.AddCell()
	cell.Value = "File_url"
	for i,v:=range orList{
		row=rows[i]
		row.SetHeightCM(1)
		cell = row.AddCell()
		cell.Value = strconv.Itoa(int(v.ID))
		cell = row.AddCell()
		cell.Value = v.CreatedAt.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.Value =v.UpdatedAt.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.Value =""
		cell = row.AddCell()
		cell.Value = v.Order_no
		cell = row.AddCell()
		cell.Value = v.User_name
		cell = row.AddCell()
		cell.Value = strconv.Itoa(int(v.Amount))
		cell = row.AddCell()
		cell.Value =v.Status
		cell = row.AddCell()
		cell.Value =v.File_url
	}
	err = file.Save("static/order.xlsx")
	if err != nil {
		return err
	}
	return nil
}
/*func Test_Add_1(t *testing.T) {
	orList:=SelA()
	if orList!=nil{
		fmt.Println(orList)
	}else{
		fmt.Println("it's wrong")
	}
}*/
