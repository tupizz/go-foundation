package main

import (
	"fmt"
	"github.com/tupizz/go-foundation/41-api-project/configs"
)

func main() {
	cfg := configs.LoadConfig(".")
	fmt.Println(cfg.DBName)

}
