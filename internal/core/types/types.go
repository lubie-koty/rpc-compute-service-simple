package types

type Server interface {
	Serve() error
}

type MathService interface {
	Add(a, b float32) float32
	Sub(a, b float32) float32
	Mul(a, b float32) float32
	Div(a, b float32) float32
}
