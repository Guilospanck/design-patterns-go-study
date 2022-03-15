package implementations

import (
	"fmt"
	"sync"
)

// The sync.Once will only perform the operation once
var once sync.Once

type single struct{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance...")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Instance already created!")
	}

	return singleInstance
}
