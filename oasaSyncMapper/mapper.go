package oasaSyncMapper

import (
	"reflect"

	"github.com/fatih/structs"
	"github.com/oulisnikos/oasaLibDb/oasaSyncModel"
	"github.com/oulisnikos/oasaLibDb/oasaSyncUtils"
)

// With this Mapper map the records from the Oasa Server to structures that we have defined
// to implement procedures for the needs of the application
func internal_mapper(source map[string]interface{}, target interface{}) {
	rvTarget := reflect.ValueOf(target)
	trvTarget := reflect.TypeOf(target)

	if rvTarget.Kind() == reflect.Pointer {
		rvTarget = rvTarget.Elem()
		trvTarget = trvTarget.Elem()
		target = reflect.New(rvTarget.Type())
	}
	for i := 0; i < rvTarget.NumField(); i++ {
		field := rvTarget.Field(i)
		fieldType := field.Kind()
		v := rvTarget.Field(i)
		tag, _ := trvTarget.Field(i).Tag.Lookup("json")
		if len(tag) != 0 {
			// v.Set(reflect.ValueOf(source[tag]))
			sourceFieldVal := source[tag]
			if sourceFieldVal != nil {
				switch fieldType {
				case reflect.String:
					v.SetString(sourceFieldVal.(string))
				case reflect.Int64:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToInt64(sourceFieldVal)))
				case reflect.Int32:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToInt32(sourceFieldVal)))
				case reflect.Int16:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToInt16(sourceFieldVal)))
				case reflect.Int8:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToInt8(sourceFieldVal)))
				case reflect.Float32:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToFloat32(sourceFieldVal)))
				case reflect.Float64:
					v.Set(reflect.ValueOf(oasaSyncUtils.StrToFloat(sourceFieldVal)))
				case reflect.Ptr:
					v.Set(reflect.ValueOf(nil))
				}
			}
		}
	}
}

// Function to Map structures from one to another with same field data types
// but one of them has less fields from the other
func structMapper(source interface{}, target interface{}) {
	sourceMap := structs.Map(source)
	rvTarget := reflect.ValueOf(target)
	trvTarget := reflect.TypeOf(target)

	if rvTarget.Kind() == reflect.Pointer {
		rvTarget = rvTarget.Elem()
		trvTarget = trvTarget.Elem()
		target = reflect.New(rvTarget.Type())
	}
	for i := 0; i < rvTarget.NumField(); i++ {
		v := rvTarget.Field(i)
		// v.Set(reflect.ValueOf(source[tag]))
		sourceFieldVal := sourceMap[trvTarget.Field(i).Name]
		if sourceFieldVal != nil {
			v.Set(reflect.ValueOf(sourceFieldVal))
		}
	}
}

func BussLineMapper(source map[string]interface{}) oasaSyncModel.Busline {
	var busLineOb oasaSyncModel.Busline
	internal_mapper(source, &busLineOb)

	return busLineOb
}

func BusRouteMapper(source map[string]interface{}) oasaSyncModel.BusRoute {
	var busRouteOb oasaSyncModel.BusRoute
	internal_mapper(source, &busRouteOb)

	return busRouteOb
}

func BusStopDtoMapper(source map[string]interface{}) oasaSyncModel.BusStopDto {
	var busStopOb oasaSyncModel.BusStopDto
	internal_mapper(source, &busStopOb)
	return busStopOb
}

func BusStopDtoToBusStop(source oasaSyncModel.BusStopDto) oasaSyncModel.BusStop {
	var busStop oasaSyncModel.BusStop
	structMapper(source, &busStop)
	return busStop
}

func SheduleMasterLineDtoMapper(source map[string]interface{}) oasaSyncModel.BusScheduleMasterLineDto {
	var busSheduleMaster oasaSyncModel.BusScheduleMasterLineDto
	internal_mapper(source, &busSheduleMaster)
	return busSheduleMaster
}

func ScheduleMasterLineDtoToScheduleMasterLine(source oasaSyncModel.BusScheduleMasterLineDto) oasaSyncModel.BusScheduleMasterLine {
	var result oasaSyncModel.BusScheduleMasterLine
	structMapper(source, &result)
	return result
}
