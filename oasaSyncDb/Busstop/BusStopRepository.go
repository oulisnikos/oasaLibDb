package Busstop

import (
	"fmt"

	"github.com/oulisnikos/oasaLibDb/logger"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
	"gorm.io/gorm"
)

func SelectByStopCode(stopCode int64) (*oasaSyncModel.BusStop, error) {
	var selectedVal oasaSyncModel.BusStop
	r := oasaSyncDb.DB.Table("BUSSTOP").Where("stop_code = ?", stopCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func Save(busStop oasaSyncModel.BusStop) error {
	selectedBusStop, err := SelectByStopCode(busStop.Stop_code)
	if err != nil {
		return err
	}
	isNew := selectedBusStop == nil
	var r *gorm.DB = nil
	if isNew {
		logger.INFO(fmt.Sprintf("Bus Stop not found [stop_code: %d]. Create New.\n", busStop.Stop_code))
		busStop.Id = oasaSyncDb.SequenceGetNextVal(oasaSyncModel.BUSSTOP_SEQ)
		//input.Line_descr = input.Line_descr + " New"

		r = oasaSyncDb.DB.Table("BUSSTOP").Create(&busStop)

	} else {
		logger.INFO(fmt.Sprintf("Bus Stop [stop_code: %d]. Updated.\n", busStop))
		busStop.Id = selectedBusStop.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = oasaSyncDb.DB.Table("BUSSTOP").Save(&busStop)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func StopList01(routeCode int32) ([]oasaSyncModel.StopDto, error) {
	var result []oasaSyncModel.StopDto
	r := oasaSyncDb.DB.Table("BUSSTOP").
		Select("BUSSTOP.*, "+
			"BUSROUTESTOPS.senu").
		Joins("LEFT JOIN BUSROUTESTOPS ON BUSSTOP.stop_code=BUSROUTESTOPS.stop_code").
		Where("BUSROUTESTOPS.route_code=?", routeCode).Order("senu").Find(&result)
	if r.Error != nil {
		logger.ERROR(r.Error.Error())
		return nil, r.Error
	}
	return result, nil
}
