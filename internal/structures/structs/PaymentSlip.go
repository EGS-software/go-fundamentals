package structs

import (
	"setup/internal/structures/interfaces"
)

type PaymentSlip struct {
	processor interfaces.Processor
	id        int
	amount    float64
}
