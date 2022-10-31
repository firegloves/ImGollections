package main

import (
	"ImGollections/config"
	"ImGollections/generator"
	"fmt"
	"log"
	"os"
)

func main() {
	log.Printf("NTGGonarator started")

	config.LoadConfig()

	fmt.Println(config.GetConfig().Layerfolders)

	err := generator.Generate(config.GetConfig().Layerfolders)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
