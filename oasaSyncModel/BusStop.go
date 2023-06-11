package oasaSyncModel

type BusStop struct {
	Id               int64       `json:"Id" gorm:"primaryKey"`
	Stop_code        int64       `json:"StopCode"`
	Stop_id          string      `json:"StopID"`
	Stop_descr       string      `json:"StopDescr"`
	Stop_descr_eng   string      `json:"StopDescrEng"`
	Stop_street      interface{} `json:"StopStreet"`
	Stop_street_eng  interface{} `json:"StopStreetEng"`
	Stop_heading     int32       `json:"StopHeading"`
	Stop_lat         float32     `json:"StopLat"`
	Stop_lng         float32     `json:"StopLng"`
	Route_stop_order int16       `json:"RouteStopOrder"`
	Stop_type        int8        `json:"StopType"`
	Stop_amea        int8        `json:"StopAmea"`
}
