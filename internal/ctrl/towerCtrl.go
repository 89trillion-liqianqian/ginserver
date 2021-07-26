package ctrl

import (
	"ginserver/internal/handler"
	"ginserver/internal/myerr"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// ping
func PingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
	return
}

// 获取士兵，根据稀有度，解锁，cvc
func GetArmyByRarity(c *gin.Context) {
	rarity := c.Query("rarity")
	unlockArena := c.Query("unlockArena")
	cvc := c.Query("cvc")
	msg := "ok 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵"
	// 参数校验
	rarityInt, _ := strconv.Atoi(rarity)
	unlockArenaInt, _ := strconv.Atoi(unlockArena)
	if rarityInt < 0 || unlockArenaInt < 0 {
		log.Println("--err GetArmyByRarity", rarity, unlockArena, cvc)
		msg = "参数错误"
		myerr.ResponseErr(c, msg)
		return
	}
	reData, err := handler.GetArmyByRarityHandler(rarityInt, unlockArenaInt, cvc)
	if err != nil {
		msg = "参数错误"
		myerr.ResponseErr(c, msg)
		return
	}
	if len(reData) < 1 {
		msg = "无符合条件的士兵"
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": reData,
	})
	return
}

//输入士兵id获取稀有度
func GetRarity(c *gin.Context) {
	armyId := c.Query("armyId")
	msg := "ok 输入士兵id获取稀有度"
	if len(armyId) < 1 {
		log.Println("--err GetRarity", armyId)
		msg = "参数错误"
		myerr.ResponseErr(c, msg)
		return
	}
	reData, err := handler.GetRarityHandler(armyId)
	if len(reData) < 1 || err != nil {
		msg = "输入士兵id不存在"
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": reData,
	})
	return
}

//输入士兵id获取战力
func GetCombatPoints(c *gin.Context) {
	armyId := c.Query("armyId")
	msg := "ok 输入士兵id获取战力"
	if len(armyId) < 1 {
		log.Println("--err GetCombatPoints", armyId)
		msg = "参数错误"
		myerr.ResponseErr(c, msg)
		return
	}
	reData, err := handler.GetCombatPointsHandler(armyId)
	if len(reData) < 1 || err != nil {
		msg = "输入士兵id不存在"
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "输入士兵id获取战力",
		"data": reData,
	})
	return
}

//输入cvc获取所有合法的士兵
func GetArmyByCvc(c *gin.Context) {
	cvc := c.Query("cvc")
	msg := "ok 输入cvc获取所有合法的士兵"
	reData, err := handler.GetArmyByCvcHandler(cvc)
	if len(reData) < 1 || err != nil {
		msg = "输入cvc不存在"
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": reData,
	})
	return
}

//获取每个阶段解锁相应士兵的json数据
func GetArmyGroupUnlockArena(c *gin.Context) {
	msg := "ok 获取每个阶段解锁相应士兵的json数据"
	// 获取每个阶段的配置
	reData, err := handler.GetArmyGroupUnlockArenaHandler()
	if len(reData) < 1 || err != nil {
		msg = "配置不存在"
		myerr.ResponseErr(c, msg)
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  msg,
		"data": reData,
	})
	return
}
