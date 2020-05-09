package bill

import (
	"time"
)

type Payment struct {
	ID             string
	ContractId     string
	CreationDate   time.Time
	ExpirationDate time.Time
	Payed          bool
}
