package structs

import (
	"setup/internal/structures/interfaces"
)

type CalcApp struct {
	processor interfaces.Processor
	id        int
	name      string
}
