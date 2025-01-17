package main

import (
	"RESTAPIporject/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)
}
