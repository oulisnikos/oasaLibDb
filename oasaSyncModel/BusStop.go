package oasaSyncModel

type BusStop struct {
	Id               int64   `json:"Id" gorm:"primaryKey"`
	Route_code       int64   `json:"RouteCode" gorm:"index:ROUTE_CODE_INDX"`
	Stop_code        int64   `json:"StopCode" gorm:"index:STOP_CODE_UN,unique"`
	Stop_id          string  `json:"StopID"`
	Stop_descr       string  `json:"StopDescr"`
	Stop_descr_eng   string  `json:"StopDescrEng"`
	Stop_street      string  `json:"StopStreet"`
	Stop_street_eng  string  `json:"StopStreetEng"`
	Stop_heading     int32   `json:"StopHeading"`
	Stop_lat         float64 `json:"StopLat"`
	Stop_lng         float64 `json:"StopLng"`
	Route_stop_order int16   `json:"RouteStopOrder"`
	Stop_type        int8    `json:"StopType"`
	Stop_amea        int8    `json:"StopAmea"`
}
