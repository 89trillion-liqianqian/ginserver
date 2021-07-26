## 1.整体框架

整理士兵配置文件格式且只保留有用的数据，使用Gin开发一个服务。

1） 启动时读取app.ini 配置文件，监听配置中的http端口号，用户请求过来时，返回有用的士兵配置。

2） 启动时用github.com/spf13/pflag 来解析命令行参数传入的 config.army.model.json文件路径，整理格式且只保留有用的数据，生成新的json文件。启动时解析该文件。

仅部分下列逻辑中的士兵信息有用，其他部分无用，设计合适的数据结构处理数据并保存，并封装如下方法：

1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵 

2）输入士兵id获取稀有度

3）输入士兵id获取战力

4）输入cvc获取所有合法的士兵 

5）获取每个阶段解锁相应士兵的json数据

## 2.目录结构

```
目录：
liqianqian@liqianqian ginserver % pwd
/Users/liqianqian/go/src/ginserver
项目结构分析：
liqianqian@liqianqian ginserver % tree
├── README.md											# 技术文档												
├── app
│   ├── http
│   │   └── httpServer.go
│   └── main.go
├── config											# 服务配置文件
│   ├── app.ini
│   ├── config.army.model.json
│   └── config.json											#处理后的士兵配置
├── go.mod
├── go.sum
├── internal
│   ├── ctrl											#控制器
│   │   └── towerCtrl.go
│   ├── handler
│   │   ├── config.json
│   │   ├── towerHandler.go
│   │   └── towerHandler_test.go											#单元测试
│   ├── model
│   │   └── config.go											#model 配置解析
│   ├── myerr											#错误返回
│   │   └── myerr.go
│   ├── router											#路由
│   │   └── router.go
│   └── service
│       └── service.go
├── locust											#压测
│   ├── __pycache__
│   │   ├── load.cpython-37.pyc
│   │   └── locust.cpython-37.pyc
│   ├── load.py
│   └── report_1626786190.391604.html
├── test
│   └── test.go
└── 题一流程图.jpg											#流程图

```

## 3.逻辑代码分层

|    层     | 文件夹                            | 主要职责        | 调用关系                  | 其它说明     |
| :-------: | :-------------------------------- | --------------- | ------------------------- | ------------ |
|  应用层   | /app/http/httpServer.go           | http 服务器启动 | 调用路由层                | 不可同层调用 |
|  路由层   | /internal/router/router.go        | 路由转发        | 被应用层调用，调用控制层  | 不可同层调用 |
|  控制层   | /internal/ctrl/tower.go           | 士兵配置获取    | 被路由层调用，调用handler | 不可同层调用 |
| handler层 | /internal/handler/towerHandler.go | 处理具体业务    | 被控制层调用              | 不可同层调   |
|   model   | /internal/model                   | config配置加载  | 被控制层调用              |              |
| 压力测试  | Locust/load.py                    | 进行压力测试    | 无调用关系                | 不可同层调用 |

## 4.存储设计

无

## 5.接口设计供客户端调用的接口

5.1输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵

请求方法

http get 

接口地址：

127.0.0.1:8000/getArmyByRarity

请求参数：

```
{
	"rarity":1,//稀有度
	"unlockArena"=1,//解锁阶段
	"cvc":1000,//客户端版本
}
```

json

请求响应

```
{
	"code":0,
	"msg":""ok,
	"data":[
		{
		"id":10101,			// 士兵ID
		"note":"步兵lv1",	// 名称
		...
		}
	]
}
```

响应状态码

| 状态码 | 说明            |
| ------ | --------------- |
| 0      | 请求成功        |
| 1      | rarity 不存在   |
| 2      | 解锁阶段 不存在 |
| 3      | cvc  不存在     |

5.2输入士兵id获取稀有度

请求方法

http get 

接口地址：

127.0.0.1:8000/getRarity

请求参数：

```
{
	"armyId":10010,//士兵id
}
```

json

请求响应

```
{
	"code":0,
	"msg":""ok,
	"data":{
		"id":10101,	// 士兵ID
		"rarity":2	// 稀有度
		}
}
```

响应状态码

| 状态码 | 说明                |
| ------ | ------------------- |
| 0      | 请求成功            |
| 1      | armyId士兵ID 不存在 |

5.3输入士兵id获取战力

请求方法

http get 

接口地址：

127.0.0.1:8000/getCombatPoints

请求参数：

```
{
	"armyId":10010,//士兵id
}
```

json

请求响应

```
{
	"code":0,
	"msg":""ok,
	"data":{
		"id":10101,	// 士兵ID
		"combatPoints":2	// 战力
		}
}
```

响应状态码

| 状态码 | 说明                |
| ------ | ------------------- |
| 0      | 请求成功            |
| 1      | armyId士兵ID 不存在 |

5.4输入cvc获取所有合法的士兵 

请求方法

http get 

接口地址：

127.0.0.1:8000/getArmyByCvc

请求参数：

```
{
	"cvc":1000,//客户端版本
}
```

json

请求响应

```
{
	"code":0,
	"msg":""ok,
	"data":[
		{
		"id":10101,		// 士兵ID
		"note":"步兵lv1",	// 名称
		...
		}
	]
}
```

响应状态码

| 状态码 | 说明                |
| ------ | ------------------- |
| 0      | 请求成功            |
| 1      | cvc客户端版本不存在 |

5.5获取每个阶段解锁相应士兵的json数据

请求方法

http get 

接口地址：

127.0.0.1:8000/getArmyGroupUnlockArena

请求参数：

```
{
	无
}
```

json

请求响应

```
json
{
	"code":0,
	"msg":""ok,
	"data":{
			1:[    // key 解锁值
				{
					"id":10101,
					"note":"步兵lv1",
					...
				}
			]
	}
}
```

响应状态码

| 状态码 | 说明           |
| ------ | -------------- |
| 0      | 请求成功       |
| 1      | 士兵配置不存在 |

## 6.第三方库

gin

```
用于api服务，go web 框架
代码： github.com/gin-gonic/gin
```

pflag 

```
用于解析命令行参数
代码：github.com/spf13/pflag 
```

goconfig

```
用于解析json配置
代码:github.com/Unknwon/goconfig
```

## 7.如何编译执行

```
#切换主目录下
cd ./app/
#编译
go build
```

## 8.todo 

```
后续配置修改
```

