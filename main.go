package main

import (
	"dimo-backend/driver"
	"fmt"
)

func main() {
	db := driver.ConnectDefault()
	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connection ok")
	result, _ := driver.FactorizationRequest(131, []int64{4, 24, 49, 124, 144, 482, 394, 254, 284})
	fmt.Println(result)
	result, _ = driver.SequenceRequest([]int64{1, 234, 2, 38}, []int64{4, 24, 49, 124, 144, 482, 394, 254, 284})
	fmt.Println(result)
}