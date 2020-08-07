package dbb

import (
	"db/createTable.go/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)




func CreateB(db *gorm.DB,or model.Order) (err error){

	db.AutoMigrate(&model.Order{})

	// Migrate the schema
	//db.AutoMigrate(&model.Order{})
	//db.SingularTable(true)
	err=db.Create(&or).Error
	if err!=nil{
		return err

	}else {
		return nil
	}

	// 查询
	//var u = new(model.Userllllll)
	//db.First(u)
	//fmt.Printf("%#v\n", u)

	//var uu model.Userllllll
	//db.Find(&uu, "sex=?", 1)
	//fmt.Printf("%#v\n", uu)

	// 更新
	//db.Model(&u).Update("name","yangyuquan" )
	// 删除



}
