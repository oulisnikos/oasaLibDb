package Busline

import (
	"fmt"
	"oasaLibDb/oasaSyncDb"
	"oasaLibDb/oasaSyncModel"
)

const erroMessageTemplate = "Field validation for [%s] failed on the [%s] tag"

type OpswValidateError struct {
	Key     string
	Message string
}

func SelectByLineCode(id int64) *oasaSyncModel.Busline {
	//var selectedPtr *oasaSyncModel.Busline
	var selectedVal oasaSyncModel.Busline
	r := oasaSyncDb.DB.Table("BUSLINE").Where("line_code = ?", id).Find(&selectedVal)
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

func SaveBusLine(input oasaSyncModel.Busline) {
	selectedBusLine := SelectByLineCode(int64(input.Line_code))
	isNew := selectedBusLine == nil
	if isNew {
		input.Id = oasaSyncDb.SequenceGetNextVal(oasaSyncModel.BUSLINE_SEQ)
		//input.Line_descr = input.Line_descr + " New"
		r := oasaSyncDb.DB.Table("BUSLINE").Create(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r := oasaSyncDb.DB.Table("BUSLINE").Save(&input)
		if r.Error != nil {
			fmt.Println(r.Error.Error())
		}
	}

}

func BuslineList01() []oasaSyncModel.Busline {
	var result []oasaSyncModel.Busline
	r := oasaSyncDb.DB.Table("BUSLINE").Order("line_id").Find(&result)
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
	return result
}
