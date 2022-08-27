package facade

import (
	"fmt"
	"log"
)

/*
外观是一种结构型设计模式，能为复杂系统、程序库或框架提供一个简单（但有限）的接口。
尽管外观模式降低了程序的整体复杂度，但它同时也有助于将不需要的依赖移动到同一个位置。

外观模式会与多种组件进一步地进行沟通，而又不会向客户端暴露其内部的复杂性。
*/

func execute() {
	fmt.Println()
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
