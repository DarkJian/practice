package api

import "hex/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DBPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (adr Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := adr.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	if err = adr.db.Save(answer, `addiction`); err != nil {
		return 0, err
	}

	return answer, nil
}
func (adr Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := adr.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	if err = adr.db.Save(answer, `subtraction`); err != nil {
		return 0, err
	}

	return answer, nil
}
func (adr Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := adr.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	if err = adr.db.Save(answer, `multiplication`); err != nil {
		return 0, err
	}

	return answer, nil
}
func (adr Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := adr.arith.Division(a, b)

	if err != nil {
		return 0, err
	}

	if err = adr.db.Save(answer, `division`); err != nil {
		return 0, err
	}

	return answer, nil
}
