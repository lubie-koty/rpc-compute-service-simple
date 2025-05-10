package types

type Server interface {
	Serve() error
}

type MathService interface {
	Add(a, b float64) float64
	Sub(a, b float64) float64
	Mul(a, b float64) float64
	Div(a, b float64) float64
}
