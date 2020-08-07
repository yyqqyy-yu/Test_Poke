package dbb

import (
	"db/createTable.go/model"
	"github.com/jinzhu/gorm"
)

func SelA(db *gorm.DB)(orList []model.Order){
	db.AutoMigrate(&model.Order{})
	if err:=db.Find(&orList).Error;err!=nil{
		return orList
	}else{
		return orList
	}
}
func SelByCT(db *gorm.DB)(orList []model.Order,err error){
	db.AutoMigrate(&model.Order{})
	db.Model(&model.Order{}).Order("created_at ASC").Find(&orList)

	if err:=db.Find(&orList).Error;err!=nil{
		return orList,err
	}else{
		return orList,nil
	}
}
func SelByMY(db *gorm.DB)(orList []model.Order,err error){
	db.AutoMigrate(&model.Order{})
	db.Model(&model.Order{}).Order("amount ASC").Find(&orList)

	if err:=db.Find(&orList).Error;err!=nil{
		return orList,err
	}else{
		return orList,nil
	}
}
func SelLike(db *gorm.DB,user_name string)(orList []model.Order){


	 db.Where("user_name like ?","%"+user_name+"%").Find(&orList)
	if orList!=nil{
		return orList
	}else{
		return orList
	}
}

func SelO(db *gorm.DB,id uint,or model.Order)(err error){
	if err=  db.Where("id=?",id).First(&or).Error;err!=nil{
		return err
	}
	return nil

}

func SelOne(db *gorm.DB,id uint)(err error,bor model.Order){
	if err=  db.Where("id=?",id).First(&bor).Error;err!=nil{
		return err,bor
	}
	return nil,bor

}
