package types

type Server interface {
	Serve() error
}

type MathService interface {
	Add(a, b int32) int32
	Sub(a, b int32) int32
	Mul(a, b int32) int32
	Div(a, b int32) int32
}
