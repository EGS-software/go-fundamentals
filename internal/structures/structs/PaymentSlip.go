package structs

import (
	"fmt"
)

type PaymentSlip struct {
	id     int
	amount float64
}

func (p PaymentSlip) Process(Data string) string {
	fmt.Print("Starting Payment Slip Processing...\n")
	return fmt.Sprintf("Payment Slip Processed: %s", Data)
}

func (p PaymentSlip) CalcData(Value float64) float64 {
	return Value * 1.05 // Example: adding 5% processing fee
}
