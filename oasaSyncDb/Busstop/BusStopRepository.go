package Busstop

import (
	"fmt"
	"github.com/oulisnikos/oasaLibDb/oasaSyncDb"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
)

func SelectByStopCode(stopCode int64) *oasaSyncModel.BusStop {
	var selectedVal oasaSyncModel.BusStop
	r := oasaSyncDb.DB.Table("BUSSTOP").Where("stop_code = ?", stopCode).Find(&selectedVal)
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

func Save(input oasaSyncModel.BusStop) {
	selectedBusStop := SelectByStopCode(input.Stop_code)
	isNew := selectedBusStop == nil
	if isNew {
		input.Id = oasaSyncDb.SequenceGetNextVal(oasaSyncModel.BUSSTOP_SEQ)
		//input.Line_descr = input.Line_descr + " New"
		r := oasaSyncDb.DB.Table("BUSSTOP").Create(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}

	} else {
		input.Id = selectedBusStop.Id
		//input.Line_descr = input.Line_descr + " Update"
		r := oasaSyncDb.DB.Table("BUSSTOP").Save(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}
	}

}
