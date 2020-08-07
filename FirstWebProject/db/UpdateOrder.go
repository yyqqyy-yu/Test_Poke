package dbb

import (
	"db/createTable.go/model"
	"github.com/jinzhu/gorm"
)

func Upd(db *gorm.DB,or model.Order)(err error){
	db.AutoMigrate(&model.Order{})
	/*if err=  DB.Where("id=?",id).First(&todo).Error;err!=nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
		return
	}*/
	if err =db.Save(&or).Error;err!=nil{
		return err
	}else {
		return nil
	}
}
