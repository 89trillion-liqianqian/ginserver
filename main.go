package main

import (
	"ginserver/controller"
	"ginserver/model"
	"log"
	"strings"
)

import (
	"github.com/gin-gonic/gin"
	flag "github.com/spf13/pflag"
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

func main() {
	log.Println("-start gin server")
	// 解析命令行参数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)
	flag.Parse()
	// 获取配置文件路径，加载士兵配置
	log.Println("---config", *configPath)
	model.LoadConfigJson(*configPath)
	//model.WriteConfig()
	// 获取server port
	filepath := "app.ini"
	model.GetAppIni(filepath)
	httpPort, _ := model.GetAppPort()
	log.Println("--httpPort", httpPort)

	//由于是外部调用包，所以必须含包名 gin. 作为前缀
	//Default 返回带有已连接 Logger 和 Recovery 中间件的 Engine 实例。
	r := gin.Default()
	// Engine 结构体中内嵌了 RouterGroup 结构体，即继承了 RouterGroup（其有成员方法 GET、POST、DELETE、PUT、ANY 等）
	controller.Routers(r)
	// 默认是 0.0.0.0:8080 端口，内部使用了 http.ListenAndServe(address, engine)
	r.Run("0.0.0.0:" + httpPort) // listen and serve on 0.0.0.0:9090
}
