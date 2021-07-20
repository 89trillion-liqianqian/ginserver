package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

/**
api 测试
*/

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容

func httpGet(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

func main() {
	//ipStr:= "http://127.0.0.1:8000/getArmyByRarity?rarity=1&unlockArena=1&cvc=1000"
	ipStr := "http://127.0.0.1:8000/"
	api := ""
	result := ""
	//1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
	api = "getArmyByRarity?rarity=1&unlockArena=1&cvc=1000"
	result = httpGet(ipStr + api)
	log.Println("--输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵", result)

	//2）输入士兵id获取稀有度
	api = "getRarity?armyId=10101"
	result = httpGet(ipStr + api)
	log.Println("--输入士兵id获取稀有度", result)

	//3）输入士兵id获取战力
	api = "getCombatPoints?armyId=10101"
	result = httpGet(ipStr + api)
	log.Println("--输入士兵id获取战力", result)

	//4）输入cvc获取所有合法的士兵
	api = "getArmyByCvc?cvc=1000"
	result = httpGet(ipStr + api)
	log.Println("--输入cvc获取所有合法的士兵", result)

	//5）获取每个阶段解锁相应士兵的json数据
	api = "getArmyGroupUnlockArena"
	result = httpGet(ipStr + api)
	log.Println("--获取每个阶段解锁相应士兵的json数据", result)
}
