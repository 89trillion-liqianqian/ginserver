package handler

import (
	"ginserver/internal/model"
	"ginserver/internal/service"
)

// 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func GetArmyByRarityHandler(rarity, unlockArena int, cvc string) (reData []model.ArmyConfig, err error) {
	//获取配置
	config := model.GetConfig()
	for id := range *config {
		cData := (*config)[id]
		if rarity == cData.Rarity && cvc == cData.Cvc && unlockArena <= cData.UnlockArena {
			reData = append(reData, cData)
		}
	}
	return
}

// 输入士兵id获取稀有度
func GetRarityHandler(armyId string) (reData map[string]int, err error) {
	reData = make(map[string]int)
	//获取配置
	armyIdInt := 0
	config := model.GetConfig()
	if _, ok := (*config)[armyId]; ok {
		armyIdInt, err = service.ArmyIdAtoi(armyId)
		reData["id"] = armyIdInt
		reData["rarity"] = (*config)[armyId].Rarity
	}
	return
}

// 输入士兵id获取战力
func GetCombatPointsHandler(armyId string) (reData map[string]int, err error) {
	reData = make(map[string]int)
	//获取配置
	armyIdInt := 0
	config := model.GetConfig()
	if _, ok := (*config)[armyId]; ok {
		armyIdInt, err = service.ArmyIdAtoi(armyId)
		reData["id"] = armyIdInt
		reData["combatPoints"] = (*config)[armyId].CombatPoints
	}

	return
}

//输入cvc获取所有合法的士兵
func GetArmyByCvcHandler(cvc string) (reData []model.ArmyConfig, err error) {
	//获取配置
	config := model.GetConfig()
	for id := range *config {
		cData := (*config)[id]
		if cvc == cData.Cvc {
			reData = append(reData, cData)
		}
	}
	return
}

//获取每个阶段解锁相应士兵的json数据
func GetArmyGroupUnlockArenaHandler() (reData map[int][]model.ArmyConfig, err error) {
	reData = make(map[int][]model.ArmyConfig, 0)
	//获取配置
	config := model.GetConfig()
	for id := range *config {
		cData := (*config)[id]
		unlock := cData.UnlockArena
		if _, ok := reData[unlock]; ok {
			reData[unlock] = append(reData[unlock], cData)
		} else {
			aData := make([]model.ArmyConfig, 0)
			aData = append(aData, cData)
			reData[unlock] = aData
		}
	}
	return
}
