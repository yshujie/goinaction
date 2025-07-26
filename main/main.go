package main

import (
	"fmt"

	"github.com/yshujie/goinaction/base"
)

func init() {
	fmt.Println("in main init")
}

func main() {
	fmt.Println("... start ...")

	base.ContinuousPrint6(5, 10)

	fmt.Println("... end ...")
}
