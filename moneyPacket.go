/*
拆的时候会实时计算金额，其金额为1分到剩余平均值2倍之间随机数，一个总金 额为M元的红包，最大的红包为 M * 2 /N（且不会超过M），当拆了红包后会更新剩余金额和个数。

*/

package moneyPacket

import (
	"math"
	"math/rand"
	"time"
)

type MoneyPackage struct {
	Money float64
	Size  int32
}

func GetRandMoney(moneyPackage *MoneyPackage) float64 {
	// moneyPackage.Money *= 100
	if moneyPackage.Size == 1 {
		moneyPackage.Size--
		return moneyPackage.Money
	}

	var min float64 = 0.01
	max := (moneyPackage.Money / float64(moneyPackage.Size)) * 2

	money := math.Floor((randNum() * max * 100)) / 100

	if money <= min {
		money = min
	}

	moneyPackage.Size--
	moneyPackage.Money -= money

	//避免出现精度的问题
	moneyPackage.Money = math.Floor(moneyPackage.Money*100) / 100

	return money
}

//获取随机数
func randNum() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}
