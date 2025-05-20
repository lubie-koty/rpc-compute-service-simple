package base

type SimpleMathService struct{}

func NewSimpleMathService() *SimpleMathService {
	return &SimpleMathService{}
}

func (s *SimpleMathService) Add(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := 0.0
	for _, num := range numbers {
		result += num
	}
	return result
}

func (s *SimpleMathService) Sub(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		result -= num
	}
	return result
}

func (s *SimpleMathService) Mul(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		result *= num
	}
	return result
}

func (s *SimpleMathService) Div(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	result := numbers[0]
	for _, num := range numbers[1:] {
		if num == 0 {
			return 0
		}
		result /= num
	}
	return result
}
