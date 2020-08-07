package handler

import (
	"db/createTable.go/model"
	"db/createTable.go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)
var err error
func Ins(c *gin.Context){
	var or model.Order
	c.BindJSON(&or)
	err=service.InsO(or)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {
		c.JSON(http.StatusOK,or)

	}
}

func InsN(c *gin.Context){
	var or model.Order
	c.BindJSON(&or)
	err=service.InsN(or)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {
		c.JSON(http.StatusOK,or)
	}


}

func Upd(c *gin.Context){
	var or model.Order

	id,_:=strconv.Atoi(c.Query("id"))
	 if service.SelO(uint(id),or)!=nil{
		 c.JSON(http.StatusOK,gin.H{"error":"id不存在"})
		 	return
	 }
	c.BindJSON(&or)

	err:=service.UpdO(or)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {
		c.JSON(http.StatusOK,or)
	}

	//c.BindJSON(&or)
	//id,ok:=c.Params.Get("id")
	//if !ok{
	//	c.JSON(http.StatusOK,gin.H{"error":"id不存在"})
	//	return
//	}
	//id1,_:=strconv.Atoi(id)
/*fmt.Println(or)
	err=service.UpdO(or.ID,or)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {
		c.JSON(http.StatusOK,or)

	}*/
}
func Del(c *gin.Context){
	id,_:=strconv.Atoi(c.Query("id"))

	err:=service.DelO(uint(id))
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})

	}else {
		c.JSON(http.StatusOK,nil)

	}
}

func SelA(c *gin.Context){
	var orList []model.Order
	orList=service.SelA()
	fmt.Println(orList)
	if orList!=nil{

		c.JSON(http.StatusOK,orList)

	}else {
		c.JSON(http.StatusOK,gin.H{"error":"isEmtpy"})

	}
}

func SelO(c *gin.Context){
	var or model.Order
	id,_:=strconv.Atoi(c.Query("id"))
	err,or:=service.SelOne(uint(id))
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,or)
	}

}


func SelByCT(c *gin.Context){


	orList,err:=service.SelByCT()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,orList)
	}

}
func SelByMY(c *gin.Context){


	orList,err:=service.SelByMY()
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,orList)
	}

}

func SelLike(c *gin.Context){

	user_name:=c.Query("user_name")
	orList:=service.SelLike(user_name)
	if orList!=nil{
		c.JSON(http.StatusOK,orList)
	}else{
		c.JSON(http.StatusOK,"that's situation is impossible")
	}

}

func Upload(c *gin.Context) {
	//var or model.Order
	file, _ := c.FormFile("uploadfile")
	log.Println(file.Filename)
	// Upload the file to specific dst.
	_, err := os.Create("static/"+file.Filename)
	if err!=nil{
		fmt.Println(err)
		return
	}

	c.SaveUploadedFile(file,"static/"+file.Filename)
	//url:="http://"+"192.168.20.32:8080"+"/home/yyq/"+file.Filename
	//or.File_url=url
	//service.UpdO(or)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func SaveExcel (c *gin.Context){
	if service.SaveExcel()!=nil{
		c.JSON(http.StatusOK,"fail to create execl")
	}else {
		c.JSON(http.StatusOK,"success to create  execl")
	}
}
