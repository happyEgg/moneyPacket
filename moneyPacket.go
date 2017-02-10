/*
拆的时候会实时计算金额，其金额为1分到剩余平均值2倍之间随机数，一个总金 额为M元的红包，最大的红包为 M * 2 /N（且不会超过M），当拆了红包后会更新剩余金额和个数。

*/

package moneyPacket

import (
	"math/rand"
	"time"
)

type MoneyPackage struct {
	Money int32
	Size  int32
}

func GetRandMoney(moneyPackage *MoneyPackage) float64 {
	if moneyPackage.Size == 1 {
		moneyPackage.Size--
		return float64(moneyPackage.Money) / 100
	}

	var min int32 = 1
	max := (moneyPackage.Money / moneyPackage.Size) * 2

	money := (randInt() * max) / 100

	if money <= min {
		money = min
	}

	moneyPackage.Size--
	moneyPackage.Money -= int32(money)

	return float64(money) / 100
}

func randInt() int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int31n(100)
}
