package oasaSyncModel

type BusRoute struct {
	Id              int64   `gorm:"PrimaryKey"`
	Route_code      int32   `json:"RouteCode"`
	Line_code       int32   `json:"LineCode"`
	Route_descr     string  `json:"RouteDescr"`
	Route_descr_eng string  `json:"RouteDescrEng"`
	Route_type      int8    `json:"RouteType"`
	Route_distance  float32 `json:"RouteDistance"`
}
