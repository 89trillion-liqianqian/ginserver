package main

import (
	"ginserver/app/http"
	"ginserver/internal/model"
	flag "github.com/spf13/pflag"
	"log"
	"strings"
)

// 定义命令行参数对应的变量
//var configPath = flag.StringP("config", "c", "config.army.model.json", "Input Your config army")
var configPath = flag.String("config", "config.army.model.json", "Input Your config army")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

// 题一，入口
func main() {
	// 解析命令行参数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Parse()
	// 获取配置文件路径，加载士兵配置
	log.Println("---config", *configPath)
	model.LoadConfigJson(*configPath)
	// 解析 config
	filepath := "../config/app.ini"
	model.GetAppIni(filepath)
	http.HttpServer()
}
