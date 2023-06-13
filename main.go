package main

import (
	"fmt"
	"github.com/oulisnikos/oasaLibDb/oasaSyncApi"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncMapper"
)

func main() {
	oasaSyncDb.IntializeDb("user1", "user1password", nil, nil, "oasaDb")
	routeStops, error := oasaSyncApi.GetBusStops(4198)
	if error != nil {
		fmt.Println("Error Occured on Server Request!!")
	}
	fmt.Println("This is a Bus Stop Dto ", routeStops[0])
	fmt.Println("This is a Bus Stop     ", oasaSyncMapper.BusStopDtoToBusStop(routeStops[0]))
	//selected := Busstop.SelectByStopCode(10153)
	//fmt.Println(selected)
}
