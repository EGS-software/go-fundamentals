package structs

import (
	"fmt"
)

type PaymentSlip struct {
	id     int
	amount float64
}

func (p PaymentSlip) Process(data string) string {
	fmt.Print("Starting Payment Slip Processing...\n")
	return fmt.Sprintf("Payment Slip Processed: %s", data)
}
