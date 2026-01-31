package main

import (
	"encoding/json"
	"fmt"
	"setup/internal/structures/structs"
)

func main() {
	paymentSlip := structs.PaymentSlip{Id: 1, Amount: 100.0}

	jsonData, err := json.Marshal(paymentSlip)
	if err != nil {
		fmt.Println("Error to generate JSON:", err)
		return
	}

	fmt.Println("JSON:", string(jsonData))
}
