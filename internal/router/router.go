package router

import (
	"ginserver/internal/ctrl"
	"github.com/gin-gonic/gin"
)

// 路由管理
func Router(r *gin.Engine) {
	r.GET("/ping", ctrl.PingFunc)
	r.GET("/getArmyByRarity", ctrl.GetArmyByRarity)
	r.GET("/getRarity", ctrl.GetRarity)
	r.GET("/getCombatPoints", ctrl.GetCombatPoints)
	r.GET("/getArmyByCvc", ctrl.GetArmyByCvc)
	r.GET("/getArmyGroupUnlockArena", ctrl.GetArmyGroupUnlockArena)
}
