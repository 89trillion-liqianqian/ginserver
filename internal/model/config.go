package model

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

/**
app配置文件解析
*/
var cfg *goconfig.ConfigFile

// 加载config 配置
func GetAppIni(filepath string) (err error) {
	// 解析app配置
	config, err := goconfig.LoadConfigFile(filepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}
	cfg = config
	return nil
}

// 获取端口号
func GetAppPort() (HttpPort string, err error) {
	// 获取app端口
	if HttpPort, err = cfg.GetValue("server", "HttpPort"); err != nil {
		fmt.Println("配置文件中不存在types", err)
		return HttpPort, nil
	}

	return HttpPort, nil
}

/**
解析士兵配置
*/
var (
	globalconfig *map[string]GlobalConfig // 原始士兵配置数据
	armyconfig   *map[string]ArmyConfig   // 士兵配置数据
)

//士兵原始配置
type GlobalConfig struct {
	Id           string `json:"id"`
	Note         string `json:"note"`
	UnlockArena  string `json:"UnlockArena"`
	Rarity       string `json:"Rarity"`
	CombatPoints string `json:"CombatPoints"`
	Cvc          string `json:"Cvc"`
}

//士兵处理后的配置
type ArmyConfig struct {
	Id           int    `json:"id"`
	Note         string `json:"note"`
	UnlockArena  int    `json:"UnlockArena"`
	Rarity       int    `json:"Rarity"`
	CombatPoints int    `json:"CombatPoints"`
	Cvc          string `json:"Cvc"`
}

// 文件处理
func LoadConfigJson(fliepath string) error {
	var config map[string]GlobalConfig
	file, err := ioutil.ReadFile(fliepath)
	if err != nil {
		fmt.Println("配置文件读取错误,找不到配置文件", err)
		return err
	}

	if err = json.Unmarshal(file, &config); err != nil {
		fmt.Println("配置文件读取失败", err)
		return err
	}
	//log.Println("--config",config["10101"])
	globalconfig = &config
	writeConfig()
	return nil
}

// 重新生成配置文件
func writeConfig() {
	// 重写json数据
	//log.Println("-2-config",globalconfig["10101"])
	wData := make(map[string]ArmyConfig, 0)
	for id := range *globalconfig {
		configData := (*globalconfig)[id]
		idInt, _ := strconv.Atoi(configData.Id)
		unlockArenaInt, _ := strconv.Atoi(configData.UnlockArena)
		rarityInt, _ := strconv.Atoi(configData.Rarity)
		combatPointsInt, _ := strconv.Atoi(configData.CombatPoints)
		armyData := ArmyConfig{}
		armyData.Id = idInt
		armyData.Note = configData.Note
		armyData.Cvc = configData.Cvc
		armyData.UnlockArena = unlockArenaInt
		armyData.Rarity = rarityInt
		armyData.CombatPoints = combatPointsInt
		if configData.UnlockArena == "" {
			armyData.UnlockArena = -1
		}
		if configData.Rarity == "" {
			armyData.Rarity = -1
		}
		wData[id] = armyData
	}
	// 需要的数据
	armyconfig = &wData
	data, _ := json.Marshal(wData)
	fp, err := os.OpenFile("../config/config.json", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// 获取士兵配置
func GetConfig() *map[string]ArmyConfig {
	return armyconfig
}
