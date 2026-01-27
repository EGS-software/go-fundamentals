package main

import (
	"fmt"
	"setup/internal/structures/structs"
)

func main() {
	// Using the PaymentSlip struct which implements the Processor interface
	paymentSlip := structs.PaymentSlip{Id: 1, Amount: 100.0}

	calcApp := structs.CalcApp{
		Processor: paymentSlip,
		Id:        101,
		Name:      "MyCalcApp",
	}

	func returnData(value float64) string {
		resultText := calcApp.Processor.Process("Data for Payment Slip")
		calc := calcApp.Processor.CalcData(	200.0)
		return fmt.Printf("Name: "+resultText+"CalcValue: %.2f\n", calc)
	}()
}
