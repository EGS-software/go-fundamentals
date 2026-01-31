package structs

import (
	"fmt"
	"setup/internal/structures/interfaces"
	"strconv"
)

type CalcApp struct {
	Processor interfaces.Processor
	Id        int
	Name      string
}

func (c CalcApp) Calc(value float64) (string, error) {
	resultText := c.Processor.Process("Data for Payment Slip")
	calc, err := c.Processor.CalcData(value)

	if err != nil {
		return "", fmt.Errorf("erro no cálculo da app %s: %w", c.Name, err)
	}
	result := "Name: " + resultText + " | " + "CalcValue: " + strconv.FormatFloat(calc, 'f', 2, 64)
	return result, nil
}
