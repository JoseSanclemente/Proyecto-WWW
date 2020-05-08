package payment

import (
	"time"
)

type Payment struct {
	ContractId string
	Date       time.Time
}
