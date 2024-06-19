package main

import (
	"fmt"
	"restful-api/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bye!")
}
