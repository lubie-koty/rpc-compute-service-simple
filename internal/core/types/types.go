package types

type Server interface {
	Serve() error
}

type MathService interface {
	Add(numbers []float64) float64
	Sub(numbers []float64) float64
	Mul(numbers []float64) float64
	Div(numbers []float64) float64
}
