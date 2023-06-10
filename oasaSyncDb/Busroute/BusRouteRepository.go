package Busroute

import (
	"fmt"
	"oasaLibDb/oasaSyncDb"
	"oasaLibDb/oasaSyncModel"
)

func SelectByRouteCode(routeCode int32) *oasaSyncModel.BusRoute {
	var selectedVal oasaSyncModel.BusRoute
	r := oasaSyncDb.DB.Table("BUSROUTE").Where("route_code = ?", routeCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		if r.RowsAffected == 0 {
			fmt.Println("Record does not exist!!!")
			return nil
		}
	}
	return &selectedVal
}

func SelectRouteByLineCode(line_code int32) []oasaSyncModel.BusRoute {
	var selectedVal []oasaSyncModel.BusRoute
	r := oasaSyncDb.DB.Table("BUSROUTE").Where("line_code = ?", line_code).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		if r.RowsAffected == 0 {
			fmt.Println("Record does not exist!!!")
			return nil
		}
	}
	return selectedVal
}

func Save(input oasaSyncModel.BusRoute) {
	selectedBusLine := SelectByRouteCode(int32(input.Route_code))
	isNew := selectedBusLine == nil
	if isNew {
		input.Id = oasaSyncDb.SequenceGetNextVal(oasaSyncModel.BUSROUTE_SEQ)
		//input.Line_descr = input.Line_descr + " New"
		r := oasaSyncDb.DB.Table("BUSROUTE").Create(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r := oasaSyncDb.DB.Table("BUSROUTE").Save(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}
	}

}
