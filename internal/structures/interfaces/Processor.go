package interfaces

type Processor interface {
	Process(data string) string
	CalcData(value float64) (float64, error)
}
