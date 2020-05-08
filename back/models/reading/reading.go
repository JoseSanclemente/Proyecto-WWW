package reading

import (
	"time"
)

type Reading struct {
	ID         string
	ContractID string
	Value      int
	Date       time.Time
}
