package base

type SimpleMathService struct{}

func NewSimpleMathService() *SimpleMathService {
	return &SimpleMathService{}
}

func (s *SimpleMathService) Add(a, b float64) float64 {
	return a + b
}

func (s *SimpleMathService) Sub(a, b float64) float64 {
	return a - b
}

func (s *SimpleMathService) Mul(a, b float64) float64 {
	return a * b
}

func (s *SimpleMathService) Div(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
