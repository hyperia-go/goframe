package frame

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

func (elements Elements) Max() (Element, int) {
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
		if e.Val.(float64) > m.Val.(float64) {
			m = e
			index = i
		}
	}

	return m, index
}