package creational

import (
	"fmt"
	"sync"
)

type Singleton struct {
	connectionString string
}

var instance *Singleton
var once sync.Once

func GetInstance(connection string) *Singleton {
	once.Do(func() {
		instance = &Singleton{connectionString: connection}
	})
	return instance
}

func singleton() {
	firstInstance := GetInstance("connection_string")
	fmt.Println("Instance Created: ", firstInstance.connectionString)

	secondInstance := GetInstance("updated_connection_string")
	fmt.Println("Second Instance Created: ", secondInstance.connectionString)

	if firstInstance == secondInstance {
		fmt.Println("Both instances are same")
	} else {
		fmt.Println("Both instances are not same")
	}
}
