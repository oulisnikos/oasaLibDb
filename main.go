package main

import (
	"fmt"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb/Busstop"
)

func main() {
	oasaSyncDb.IntializeDb("user1", "user1password", nil, nil, "oasaDb")
	selected := Busstop.SelectByStopCode(10153)
	fmt.Println(selected)
}
