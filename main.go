package main

import (
	"fmt"
	"github.com/oulisnikos/oasaLibDb/oasaSyncApi"
)

func main() {
	result, error := oasaSyncApi.GetBusStops(4198)
	if error != nil {
		fmt.Println("An error occured ", error.Error_Descr)
		return
	}
	fmt.Println("This is the result ", result)
}
