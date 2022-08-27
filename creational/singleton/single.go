package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

// DCL: double check lock
func getInstance() *single {
	// 如果singleInstance已经初始化过，则不用加锁了，提高效率
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// 如果一开始有多个协程进入此分支，只有一个获取到锁，并成功初始化；
		// 剩下的协程等待锁释放之后，依然会走获取锁的逻辑，所以还需要判断一次；
		// 如果已经初始化过了，则不需要再次初始化
		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

var once sync.Once

// 如果用sync.Once，则不需要使用DCL
func getInstancePlus() *single {
	if singleInstance == nil {
		// sync.Once保证只执行一次
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
