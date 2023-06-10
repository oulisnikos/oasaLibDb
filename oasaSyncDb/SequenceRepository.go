package oasaSyncDb

import (
	"fmt"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
)

func SequenceGetNextVal(seq_name string) int64 {
	var nextVal int64 = -1
	var sequnece oasaSyncModel.Sequence
	r0 := DB.Table("SEQUENCES").Where("SEQ_GEN=?", seq_name).Find(&sequnece)
	if r0 != nil && r0.RowsAffected == 0 {
		sequnece = oasaSyncModel.Sequence{
			SEQ_GEN: seq_name,
		}
		nextVal = 1
	} else {
		nextVal = sequnece.SEQ_COUNT + 1
	}
	sequnece.SEQ_COUNT = nextVal
	r1 := DB.Table("SEQUENCES").Save(&sequnece)
	if r1 != nil {
		if r1.Error != nil {
			fmt.Println(r1.Error.Error())
			return -1
		}
	}
	return nextVal
}
