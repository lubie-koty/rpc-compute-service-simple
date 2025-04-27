package core

type MathService interface {
	Add(a, b int) int
	Sub(a, b int) int
	Mul(a, b int) int
	Div(a, b int) int
}

type SimpleMathService struct{}

func (s *SimpleMathService) Add(a, b int) int {
	return a + b
}

func (s *SimpleMathService) Sub(a, b int) int {
	return a - b
}

func (s *SimpleMathService) Mul(a, b int) int {
	return a * b
}

func (s *SimpleMathService) Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
