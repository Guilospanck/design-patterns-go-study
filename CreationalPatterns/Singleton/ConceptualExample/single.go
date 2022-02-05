package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		// We're verifying here again because if more than one goroutine passes the first check, it'd create more than one instance
		if singleInstance == nil {
			fmt.Println("Creating single instance.")
			singleInstance = &single{}
		} else {
			fmt.Println("Instance already created!")
		}

	} else {
		fmt.Println("Instance already created!")
	}

	return singleInstance
}
