package router


import (
"db/createTable.go/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

//r filename []string
//r qkfilename []string
type Filename struct {
	FileList []string  `json:"filelist"`
}
type testJs struct {
	Name string
	Sex int
}
func Rounter()*gin.Engine {
	router := gin.Default()
	//var fe Filename
	router.Static("/static", "static")
	router.LoadHTMLFiles("htmlfile/index.html",)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/download", func(c *gin.Context) {
		var ts testJs
		ts.Name="ss"
		ts.Sex=1
		dir_list, e := ioutil.ReadDir("static/")
		if e != nil {
			fmt.Println("read dir error")
			return
		}
		fmt.Println(dir_list)
		for i, v := range dir_list {
			fmt.Println(i, "=", v.Name())
			//fe.fileList=append(fe.fileList,v.Name())

		}
		//fmt.Println(fe)
		c.JSON(http.StatusOK,ts)
		//filename=qkfilename
	})
	router.MaxMultipartMemory = 8 << 20
	router.GET("/OrderG", handler.SelA)
	router.GET("/OrderO", handler.SelO)
	router.GET("/OrderBYCT", handler.SelByCT)
	router.GET("/OrderBYMY", handler.SelByMY)
	router.GET("/OrderGBYNE", handler.SelLike)
	router.GET("/OrderGExcel", handler.SaveExcel)
	router.POST("/OrderP",handler.Ins)
	router.POST("/OrderP2",handler.InsN)
	router.PUT("/OrderU", handler.Upd)
	router.DELETE("/OrderD", handler.Del)
	router.POST("/upload",handler.Upload)
	//router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
return nil
}

