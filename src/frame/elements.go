package frame

import "fmt"

type Elements struct {
	Vals []Element
}

// ------------------------------------------
// Elements Methods -------------------------
// ------------------------------------------
func (elements Elements) Len() int {
	return len(elements.Vals)
}

func (elements Elements) Sum() float64 {
	sum := Element{Val: float64(0)}
	for _, x := range elements.Vals {
		sum, _ = sum.Add(x)
	}
	return sum.Val.(float64)
}

func (elements Elements) Prod() float64 {
	prod := Element{Val: float64(1)}
	for _, x := range elements.Vals {
		prod, _ = prod.Prod(x)
	}
	return prod.Val.(float64)
}

func (elements Elements) minmax(op string) (Element, int) {
	switch op {
	case "max":
		op = ">"
	case "min":
		op = "<"
	default:
		p := fmt.Sprintf("Unknown operation %s", op)
		panic(p)
	}

	m := Element{Val: elements.Vals[0].Val}
	e_float := m.AsFloat()
	index := 0
	if e_float != nil {
		panic(e_float)
	}

	for i, e := range elements.Vals {
		err := e.AsFloat()
		if err != nil {
			panic(err)
		}
		if Ops[op](e.Val.(float64), m.Val.(float64)).Val.(bool) {
			m = e
			index = i
		}
	}

	return m, index
}

func (elements Elements) Max() (Element, int) {
	return elements.minmax("max")
}

func (elements Elements) Min() (Element, int) {
	return elements.minmax("min")
}