package structs

import (
	"setup/internal/structures/interfaces"
)

type CalcApp struct {
	Processor interfaces.Processor
	Id        int
	Name      string
}
