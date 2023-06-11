package oasaSyncMapper

import (
	"reflect"

	oasaSyncModel "github.com/oulisnikos/oasaPlugin/oasaSyncModel"
	oasa_sync_utils "github.com/oulisnikos/oasaPlugin/oasaSyncUtils"
)

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
			switch fieldType {
			case reflect.String:
				v.SetString(sourceFieldVal.(string))
			case reflect.Int64:
				v.Set(reflect.ValueOf(oasa_sync_utils.StrToInt64(sourceFieldVal)))
			case reflect.Int32:
				v.Set(reflect.ValueOf(oasa_sync_utils.StrToInt32(sourceFieldVal)))
			case reflect.Int16:
				v.Set(reflect.ValueOf(oasa_sync_utils.StrToInt16(sourceFieldVal)))
			case reflect.Int8:
				v.Set(reflect.ValueOf(oasa_sync_utils.StrToInt8(sourceFieldVal)))
			case reflect.Float64:
				v.Set(reflect.ValueOf(oasa_sync_utils.StrToFloat(sourceFieldVal)))
			case reflect.Ptr:
				v.Set(reflect.ValueOf(nil))
			}
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
