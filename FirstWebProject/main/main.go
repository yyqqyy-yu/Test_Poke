package main

import (
	dbb "db/createTable.go/db"
	"db/createTable.go/router"
)

func main(){
db:=dbb.Minit("db/db.yaml")
router.Rounter()
defer db.Close()
	//pwd,_ := os.Getwd()
	//获取文件或目录相关信息
	/*fileInfoList,err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name())  //打印当前文件或目录下的文件或目录名
	}*/
	/*filepath.Walk(pwd,func(path string, info os.FileInfo, err error) error{
		fmt.Println(path) //打印path信息
		fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})*/
	/*filepathNames,err := filepath.Glob(filepath.Join(pwd,"*"))
	if err != nil {
		log.Fatal(err)
	}

	for i := range filepathNames {
		fmt.Println(filepathNames[i]) //打印path
	}*/
	/*dir_list, e := ioutil.ReadDir("static/")
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dir_list {
		fmt.Println(i, "=", v.Name())
	}*/

}

