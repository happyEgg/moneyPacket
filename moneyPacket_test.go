package moneyPacket

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMoneyPacket(t *testing.T) {
	//money的值已经扩大100倍。实际是20RMB，3个人分
	leftMoneyPackage := &MoneyPackage{2000, 3}
	end := false
	var mutex sync.Mutex

	for {
		if end {
			break
		}

		go func() {
			mutex.Lock()
			if leftMoneyPackage.Size <= 0 {
				end = true
				return
			}
			money := GetRandMoney(leftMoneyPackage)
			fmt.Println(money)
			mutex.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
}
