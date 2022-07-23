package main

import (
	"fmt"

	"github.com/ramilmsh/docker-dev-test/internal/adder"
)

func main() {
	fmt.Println(adder.Add(1, 2))
}
