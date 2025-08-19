package main

import (
	"fmt"

	goroutine_test "github.com/yshujie/goinaction/goroutine"
)

func init() {
	fmt.Println("in main init")
}

func main() {
	fmt.Println("... start ...")

	goroutine_test.AlternatePrintingSingleChan()
	fmt.Println()
	goroutine_test.AlternatePrintingDualChan()
	fmt.Println()
	fmt.Println("... end ...")
}
