package controller

import (
	"ginserver/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

/**
http get post 处理
*/

func Routers(r *gin.Engine) {
	r.GET("/ping", pingFunc)
	r.GET("/getArmyByRarity", GetArmyByRarity)
	r.GET("/getRarity", GetRarity)
	r.GET("/getCombatPoints", GetCombatPoints)
	r.GET("/getArmyByCvc", GetArmyByCvc)
	r.GET("/getArmyGroupUnlockArena", GetArmyGroupUnlockArena)
	return
}

// 测试
func pingFunc(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong22",
	})
	return
}

// 获取士兵，根据稀有度，解锁，cvc

func GetArmyByRarity(c *gin.Context) {
	rarity := c.Query("rarity")
	unlockArena := c.Query("unlockArena")
	cvc := c.Query("cvc")
	rarityInt, _ := strconv.Atoi(rarity)
	unlockArenaInt, _ := strconv.Atoi(unlockArena)
	log.Println("--unlockArena", rarity, unlockArena, cvc)
	reData := model.GetArmyByRarityModel(rarityInt, unlockArenaInt, cvc)
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵",
		"data": reData,
	})
	return
}

//输入士兵id获取稀有度

func GetRarity(c *gin.Context) {
	armyId := c.Query("armyId")
	log.Println("--armyId", armyId)
	reData := model.GetRarityModel(armyId)
	if len(reData) < 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "输入士兵id不存在",
			"data": "",
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "输入士兵id获取稀有度",
		"data": reData,
	})
	return
}

//输入士兵id获取战力

func GetCombatPoints(c *gin.Context) {
	armyId := c.Query("armyId")
	log.Println("--armyId", armyId)
	reData := model.GetCombatPointsModel(armyId)
	if len(reData) < 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "输入士兵id不存在",
			"data": "",
		})

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
	log.Println("--cvc", cvc)
	reData := model.GetArmyByCvcModel(cvc)
	if len(reData) < 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "输入cvc不存在",
			"data": "",
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "输入cvc获取所有合法的士兵",
		"data": reData,
	})
	return
}

//获取每个阶段解锁相应士兵的json数据

func GetArmyGroupUnlockArena(c *gin.Context) {
	reData := model.GetArmyGroupUnlockArenaModel()
	if len(reData) < 1 {
		c.JSON(200, gin.H{
			"code": 1,
			"msg":  "士兵配置不存在",
			"data": "",
		})

		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "获取每个阶段解锁相应士兵的json数据",
		"data": reData,
	})
	return
}
