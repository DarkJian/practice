package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (adr Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (adr Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (adr Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (adr Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}
