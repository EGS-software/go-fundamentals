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

	result, err := calcApp.Calc(paymentSlip.Amount)

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("The generic result is:", result)
	}

	fmt.Println(result)
}
