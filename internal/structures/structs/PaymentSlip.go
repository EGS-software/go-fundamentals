package structs

import (
	"errors"
	"fmt"
)

type PaymentSlip struct {
	Id     int
	Amount float64
}

// Process Implementing the Processor interface methods for PaymentSlip
func (p PaymentSlip) Process(Data string) string {
	fmt.Print("Starting Payment Slip Processing...\n")
	return fmt.Sprintf("Payment Slip Processed: %s", Data)
}

// CalcData method implementation
func (p PaymentSlip) CalcData(Value float64) (float64, error) {
	if Value < 0 {
		return 0.0, errors.New("value must be greater than zero")
	}

	return Value * 1.05, nil // Example: adding 5% processing fee
}
