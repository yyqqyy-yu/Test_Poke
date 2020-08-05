package min

import (
	"fmt"
	"reflect"
)

type Kox struct {
	width string
	eee int

}


func main(){
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	i:=Kox{"ss",5}
	t := reflect.TypeOf(&i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	w := reflect.ValueOf(&i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	tag := t.Elem().Field(1).Tag  //获取定义在struct里面的标签
	name := w.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	fmt.Println("ssssssssname",name)
	fmt.Println("ssssssname",tag)
	fmt.Println("og"=="go")


}
