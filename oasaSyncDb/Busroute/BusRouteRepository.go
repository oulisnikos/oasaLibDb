package Busroute

import (
	"fmt"

	"github.com/oulisnikos/oasaLibDb/logger"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
	"gorm.io/gorm"
)

func SelectByRouteCode(routeCode int32) (*oasaSyncModel.BusRoute, error) {
	var selectedVal oasaSyncModel.BusRoute
	r := oasaSyncDb.DB.Table("BUSROUTE").Where("route_code = ?", routeCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			// fmt.Println(r.Error.Error())
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.Logger.Infof("Bus Route Not Found [route_code: %d].\n", routeCode)
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func SelectRouteByLineCode(line_code int32) []oasaSyncModel.BusRoute {
	var selectedVal []oasaSyncModel.BusRoute
	r := oasaSyncDb.DB.Table("BUSROUTE").Where("line_code = ?", line_code).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		//if r.RowsAffected == 0 {
		//	fmt.Println("Bus Routes Not Found [li: %d].")
		//	return nil
		//}
	}
	return selectedVal
}

func Save(input oasaSyncModel.BusRoute) error {
	selectedBusLine, err := SelectByRouteCode(int32(input.Route_code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		input.Id = oasaSyncDb.SequenceGetNextVal(oasaSyncModel.BUSROUTE_SEQ)
		//input.Line_descr = input.Line_descr + " New"
		r = oasaSyncDb.DB.Table("BUSROUTE").Create(&input)

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = oasaSyncDb.DB.Table("BUSROUTE").Save(&input)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil

}

func BusRouteList01() []oasaSyncModel.BusRoute {
	var result []oasaSyncModel.BusRoute
	r := oasaSyncDb.DB.Table("BUSROUTE").Order("route_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		//if r.RowsAffected == 0 {
		//	fmt.Println("Record does not exist!!!")
		//	return nil
		//}
	}
	return result
}
