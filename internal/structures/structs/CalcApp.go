package structs

import (
	"setup/internal/structures/interfaces"
	"strconv"
)

type CalcApp struct {
	Processor interfaces.Processor
	Id        int
	Name      string
}

func (c CalcApp) CalcApp(value float64) string {
	resultText := c.Processor.Process("Data for Payment Slip")
	calc := c.Processor.CalcData(value)

	result := "Name: " + resultText + " | " + "CalcValue: " + strconv.FormatFloat(calc, 'f', 2, 64)
	return result
}
