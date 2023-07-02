package BusRouteStops

import (
	"fmt"

	"github.com/oulisnikos/oasaLibDb/logger"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
)

func DeleteStopByRoute(routeCode int64) error {
	var routeStops []oasaSyncModel.BusRouteStops
	r := oasaSyncDb.DB.Table("BUSROUTESTOPS").Where("route_code=?", routeCode).Delete(&routeStops)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	logger.INFO(fmt.Sprintf("Deleted Rows are %d", r.RowsAffected))
	return nil
}

func SaveRouteStops(input oasaSyncModel.BusRouteStops) error {
	r := oasaSyncDb.DB.Table("BUSROUTESTOPS").Create(&input)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	logger.INFO(fmt.Sprintf("Stop [%d] saved Succefully in Route [%d]", input.Stop_code, input.Route_code))
	return nil
}
