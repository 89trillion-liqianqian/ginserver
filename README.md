一:gin http server 解析配置文件技术文档v1

1.目录结构

```
目录：
liqianqian@liqianqian ginserver % pwd
/Users/liqianqian/go/src/ginserver
项目结构分析：
liqianqian@liqianqian ginserver % tree
.
├── app.ini										// 服务配置文件
├── config.army.model.json		// 士兵配置
├── controller								// http api
│   └── tower.go							// 士兵 api
├── go.mod					
├── go.sum
├── main.go										// 入口函数
└── model											// 功能模块
    └── load.go								// 配置数据处理
├── README.md
├── app.ini										// 服务配置文件
├── config.army.model.json		// 士兵配置
├── config.json								// 解析后的数据
├── controller								// http api
│   └── tower.go							//士兵 api
├── go.mod
├── go.sum
├── locust										// 压测
│   ├── __pycache__
│   │   ├── load.cpython-37.pyc	
│   │   └── locust.cpython-37.pyc
│   ├── load.py							 	// 压测脚本
│   └── report_1626786190.391604.html	// 压测报告
├── main.go										// 入口函数
├── model
│   ├── load.go 							// 配置数据处理加载
│   └── requestModel.go				// model 功能模块
└── test
    └── test.go
2 directories, 7 files
liqianqian@liqianqian ginserver % 

```

2。运行

```
 示例：config 默认config.army.model.json
 go run main.go  
 或
 go run main.go --config="config.army.model.json" 
```

3.api 文档

3.1

```
1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
http get 
api: ip:port/getArmyByRarity?rarity=1&unlockArena=1&cvc=1000
请求体
rarity=稀有度&unlockArena=解锁阶段&cvc=客户端版本
响应体
json
{
	"code":0,
	"msg":""ok,
	"data":[
		{
		"id":10101,
		"note":"步兵lv1",
		...
		}
	]
}
状态码
0 ：请求成功
1 : rarity 不存在
2 : 解锁阶段 不存在
3 : cvc  不存在

```

3.2

```
2）输入士兵id获取稀有度
http get 
api: ip:port/getRarity?armyId=10101
请求体
armyId=士兵id
响应体
json
{
	"code":0,
	"msg":""ok,
	"data":{
		"id":10101,
		"rarity":2
		}
}
状态码
0 ：请求成功
1 : armyId 不存在
```

3.3

```
3）输入士兵id获取战力
http get 
api: ip:port/getCombatPoints?armyId=10101
请求体
armyId=士兵id
响应体
json
{
	"code":0,
	"msg":""ok,
	"data":{
		"id":10101,
		"combatPoints":2
		}
}
状态码
0 ：请求成功
1 : armyId 不存在

4）输入cvc获取所有合法的士兵 
5）获取每个阶段解锁相应士兵的json数据

```

3.4

```
4）输入cvc获取所有合法的士兵 
http :get
api: ip:port/getArmyByCvc?cvc=1000
请求体
cvc=客户端版本
响应体
json
{
	"code":0,
	"msg":""ok,
	"data":[
		{
		"id":10101,
		"note":"步兵lv1",
		...
		}
	]
}
状态码
0 ：请求成功
1 : cvc  不存在
```

3.5

```
5）获取每个阶段解锁相应士兵的json数据
http :get
api: ip:port/getArmyGroupUnlockArena
请求体

响应体
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
状态码
0 ：请求成功
1 ：士兵配置不存在

```

