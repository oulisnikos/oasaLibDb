package oasaSyncWeb

type OasaError struct {
	Error_Code  int32
	Error_Descr string
}

type OasaResponse struct {
	Error *OasaError
	Data  interface{}
	Retry bool
}
