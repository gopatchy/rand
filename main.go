package main

import (
	"fmt"

	"github.com/dchest/uniuri"
)

func main() {
	fmt.Printf("%s\n", uniuri.New())
}
