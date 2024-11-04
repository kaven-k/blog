package main

import (
	"fmt"
	"server/utils/random"
)

func main() {
	s := random.RandString(16)
	fmt.Println(s)
}
