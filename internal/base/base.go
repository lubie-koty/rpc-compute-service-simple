package base

type SimpleMathService struct{}

func NewSimpleMathService() *SimpleMathService {
	return &SimpleMathService{}
}

func (s *SimpleMathService) Add(a, b int32) int32 {
	return a + b
}

func (s *SimpleMathService) Sub(a, b int32) int32 {
	return a - b
}

func (s *SimpleMathService) Mul(a, b int32) int32 {
	return a * b
}

func (s *SimpleMathService) Div(a, b int32) int32 {
	if b == 0 {
		return 0
	}
	return a / b
}
