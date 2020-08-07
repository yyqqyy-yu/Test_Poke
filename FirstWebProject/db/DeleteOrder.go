package dbb

import (
	"db/createTable.go/model"
	"github.com/jinzhu/gorm"
)

func Del(db *gorm.DB,id uint)(err error){
	if err=db.Where("id=?",id).Delete(model.Order{}).Error;err!=nil {
		return err
	}else {
		return nil
	}
}