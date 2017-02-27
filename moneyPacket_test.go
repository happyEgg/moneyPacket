package moneyPacket

import (
	"fmt"
	"sync"
	"testing"
)

func TestMoneyPacket(t *testing.T) {
	leftMoneyPackage := &MoneyPackage{20.00, 3}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < int(leftMoneyPackage.Size); i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			if leftMoneyPackage.Size <= 0 {
				return
			}
			money := GetRandMoney(leftMoneyPackage)
			fmt.Println(money)
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
}
