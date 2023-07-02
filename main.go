package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/oulisnikos/oasaLibDb/logger"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb/Busstop"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	logger.InitLogger("oasaLibDb")
	oasaSyncDb.IntializeDb()
	//routeStops, error := oasaSyncApi.GetBusStops(4198)
	//if error != nil {
	//	fmt.Println("Error Occured on Server Request!!")
	//}
	//fmt.Println("This is a Bus Stop Dto ", routeStops[0])
	//busStop := oasaSyncMapper.BusStopDtoToBusStop(routeStops[0])
	//fmt.Println("This is a Bus Stop     ", busStop)
	//Busstop.Save(busStop)
	////selected := Busstop.SelectByStopCode(10153)
	////fmt.Println(selected)
	result := Busstop.StopList01(4198)
	fmt.Println(result)
}
