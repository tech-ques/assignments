package main

import (
	"common-for-all/projects/cmd/assignments/proximabiz/pkg"
	"fmt"
)

func main() {
	fmt.Println("main function started")
	start()
}

func start() {
	users, err := pkg.ReadUsersFromDisk("../test/testdata/users.json")
	if err != nil {
		fmt.Println("Error in start: while ReadUsersFromDisk, err: ", err)
	}
	// Initiate process
	pkg.StartProcess(users)
	// Initiate Retry mechanism
	pkg.RetryAll()
}
