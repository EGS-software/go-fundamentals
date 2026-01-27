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

	result := calcApp.Processor.Process("Data for Payment Slip")
	fmt.Println(result)
}
