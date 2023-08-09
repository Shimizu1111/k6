package main

import (
	"fmt"

	"go.k6.io/k6/cmd"
)

func main() {
	fmt.Println("テストのコードです")
	cmd.Execute()
}
