package service

import (
	"errors"
	"strconv"
)

// 字符串转化
func ArmyIdAtoi(str string) (value int, err error) {
	value, err = strconv.Atoi(str)
	if value < 1 {
		err = errors.New("失败id不合法")
	}
	return
}
