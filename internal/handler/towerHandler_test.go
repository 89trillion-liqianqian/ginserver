package handler

import (
	"ginserver/internal/model"
	"log"
	"testing"
)

// 初始化
func init() {
	configPath := "../../config/config.army.model.json"
	model.LoadConfigJson(configPath)
}

// 	1）输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func TestGetArmyByRarityHandler(t *testing.T) {
	rarity := 1
	unlockArena := 1
	cvc := "1000"
	reData, err := GetArmyByRarityHandler(rarity, unlockArena, cvc)
	if len(reData) < 1 || err != nil {
		log.Println("--err 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵")
		return
	}
	log.Println("ok 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵")
}

//2）输入士兵id获取稀有度
func TestGetRarityHandler(t *testing.T) {
	armyId := "10101"
	reData, err := GetRarityHandler(armyId)
	if len(reData) < 1 || err != nil {
		log.Println("--err 输入士兵id获取稀有度")
		return
	}
	log.Println("ok 输入士兵id获取稀有度")
}

//3）输入士兵id获取战力
func TestGetCombatPointsHandler(t *testing.T) {
	armyId := "10101"
	reData, err := GetCombatPointsHandler(armyId)
	if len(reData) < 1 || err != nil {
		log.Println("--err 输入士兵id获取战力")
		return
	}
	log.Println("ok 输入士兵id获取战力")
}

//4）输入cvc获取所有合法的士兵
func TestGetArmyByCvcHandler(t *testing.T) {
	cvc := "1000"
	reData, err := GetArmyByCvcHandler(cvc)
	if len(reData) < 1 || err != nil {
		log.Println("--err 输入cvc获取所有合法的士兵")
		return
	}
	log.Println("ok 输入cvc获取所有合法的士兵")
}

//5）获取每个阶段解锁相应士兵的json数据
func TestGetArmyGroupUnlockArenaHandler(t *testing.T) {
	reData, err := GetArmyGroupUnlockArenaHandler()
	if len(reData) < 1 || err != nil {
		log.Println("--err 输入cvc获取所有合法的士兵")
		return
	}
	log.Println("ok 输入cvc获取所有合法的士兵")
}
