package structs

import (
	"fmt"
)

type PaymentSlip struct {
	Id     int
	Amount float64
}

// Implementing the Processor interface methods for PaymentSlip
func (p PaymentSlip) Process(Data string) string {
	fmt.Print("Starting Payment Slip Processing...\n")
	return fmt.Sprintf("Payment Slip Processed: %s", Data)
}

// CalcData method implementation
func (p PaymentSlip) CalcData(Value float64) (float64, error) {
	return Value * 1.05 // Example: adding 5% processing fee
}
