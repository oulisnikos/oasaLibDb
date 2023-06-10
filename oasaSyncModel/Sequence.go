package oasaSyncModel

const BUSLINE_SEQ = "BUSLINE_SEQ"
const BUSROUTE_SEQ = "BUSROUTE_SEQ"

type Sequence struct {
	SEQ_GEN   string `gorm:"primaryKey"`
	SEQ_COUNT int64
}
