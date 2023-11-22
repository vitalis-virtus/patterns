package singletone

import (
	"fmt"
	"sync"
)

type Single struct{}

var lock = sync.Mutex{}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance == nil {

		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instane now!")
			singleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created!")
		}
	} else {
		fmt.Println("Single instance already created!")
	}
	return singleInstance
}
