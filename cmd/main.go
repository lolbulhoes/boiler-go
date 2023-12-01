package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lolbulhoes/boiler-go/utils"
)

func main() {
	projectName := os.Args[1]

	dir := utils.Pwd()

	err := os.Mkdir(fmt.Sprintf("%s/%s", dir, projectName), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
