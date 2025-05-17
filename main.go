package main

import (
	"ci-environment/utils"
	"fmt"
)

func main() {
	answer := utils.FindCapitalCity("Chaco")
	fmt.Println(answer)
}
