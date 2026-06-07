package main

import (
	"fmt"

	"github.com/Yeah114/gopherconvert/minecraft/world/block/full"
)

func main() {
	table := full.NewBlockRuntimeIDTable(true)
	name, properties, found := table.RuntimeIDToState(954720664)
	fmt.Println(name, properties, found)
}
