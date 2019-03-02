package element

import "fmt"

type ElementArray struct {
	Vals []Element
}

// ------------------------------------------
// ElementArray Methods -------------------------
// ------------------------------------------
func (elements ElementArray) Len() int {
	return len(elements.Vals)
}

func (elements ElementArray) Sum() float64 {
	sum := Element{Val: float64(0)}
	for _, x := range elements.Vals {
		sum, _ = Add(sum, x)
	}
	return sum.Val.(float64)
}

func (elements ElementArray) Prod() float64 {
	prod := Element{Val: float64(1)}
	for _, x := range elements.Vals {
		prod, _ = Prod(prod, x)
	}
	return prod.Val.(float64)
}

func (elements ElementArray) minmax(operation string) (Element, int) {
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

func (elements ElementArray) Max() (Element, int) {
	return elements.minmax("max")
}

func (elements ElementArray) Min() (Element, int) {
	return elements.minmax("min")
}