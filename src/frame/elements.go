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
		sum, _ = Add(sum, x)
	}
	return sum.Val.(float64)
}

func (elements Elements) Prod() float64 {
	prod := Element{Val: float64(1)}
	for _, x := range elements.Vals {
		prod, _ = Prod(prod, x)
	}
	return prod.Val.(float64)
}

func (elements Elements) minmax(operation string) (Element, int) {
	op := Le
	switch operation {
	case "max":
		op = Ge
	case "min":
		op = Le
	default:
		p := fmt.Sprintf("Unknown operation %s", op)
		panic(p)
	}

	m := elements.Vals[0]
	index := 0
	for i, e := range elements.Vals {

		res, err := op(e, m)
		if err != nil {
			panic(err)
		}
		if res.Val.(bool) {
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