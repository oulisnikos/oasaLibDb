package oasaSyncWeb

import "fmt"

type OasaError struct {
	Error_Code  int32
	Error_Descr string
}

type OasaResponse struct {
	Error *OasaError
	Data  interface{}
	Retry bool
}

func (e OasaError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Error_Code, e.Error_Descr)
}
