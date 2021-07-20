package model

import "strconv"

type ArmyByRarityModel struct {
	Rarity      int `form:"rarity" `
	UnlockArena int `form:"unlockArena"`
	Cvc         int `form:"cvc"`
}

// 输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵

func GetArmyByRarityModel(rarity, unlockArena int, cvc string) (reData []ArmyConfig) {
	//获取配置
	config := GetConfig()
	for id := range *config {
		cData := (*config)[id]
		if rarity == cData.Rarity && cvc == cData.Cvc && unlockArena <= cData.UnlockArena {
			reData = append(reData, cData)
		}
	}
	return
}

// 输入士兵id获取稀有度

func GetRarityModel(armyId string) (reData map[string]int) {
	reData = make(map[string]int)
	//获取配置
	config := GetConfig()
	if _, ok := (*config)[armyId]; ok {
		armyIdInt, _ := strconv.Atoi(armyId)
		reData["id"] = armyIdInt
		reData["rarity"] = (*config)[armyId].Rarity
	}

	return
}

// 输入士兵id获取战力

func GetCombatPointsModel(armyId string) (reData map[string]int) {
	reData = make(map[string]int)
	//获取配置
	config := GetConfig()
	if _, ok := (*config)[armyId]; ok {
		armyIdInt, _ := strconv.Atoi(armyId)
		reData["id"] = armyIdInt
		reData["combatPoints"] = (*config)[armyId].CombatPoints
	}

	return
}

//输入cvc获取所有合法的士兵

func GetArmyByCvcModel(cvc string) (reData []ArmyConfig) {
	//获取配置
	config := GetConfig()
	for id := range *config {
		cData := (*config)[id]
		if cvc == cData.Cvc {
			reData = append(reData, cData)
		}
	}
	return
}

//获取每个阶段解锁相应士兵的json数据

func GetArmyGroupUnlockArenaModel() (reData map[int][]ArmyConfig) {
	reData = make(map[int][]ArmyConfig, 0)
	//获取配置
	config := GetConfig()
	for id := range *config {
		cData := (*config)[id]
		unlock := cData.UnlockArena
		if _, ok := reData[unlock]; ok {
			reData[unlock] = append(reData[unlock], cData)
		} else {
			aData := make([]ArmyConfig, 0)
			aData = append(aData, cData)
			reData[unlock] = aData
		}
	}
	return
}
