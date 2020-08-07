package dbb

import (

	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB
type Config struct {
	Name string
}
var err error
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
func Minit(config string) (db *gorm.DB){
	if err = Init(config); err != nil {
		panic(err)
	}

	user := viper.GetString("username")
	password := viper.GetString("password")
	addr := viper.GetString("addr")
	dbName := viper.GetString("dbName")
	charset := viper.GetString("charset")
	parseTime := viper.GetString("parseTime")
	loc := viper.GetString("loc")
	fmt.Println("Viper get name:",user)
	DB, err = gorm.Open("mysql", user+":"+password+"@"+"("+addr+")"+"/"+dbName+"?charset="+charset+"&parseTime="+parseTime+"&loc="+loc)
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}