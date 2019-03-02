package frame

import "errors"

type Element struct {
	Val interface{}
}

// ------------------------------------------
// Element Methods --------------------------
// ------------------------------------------
func (e *Element) AsFloat() error {
	switch e.Val.(type) {
	case uint8:
		e.Val = float64(e.Val.(uint8))
	case int8:
		e.Val = float64(e.Val.(uint8))
	case uint16:
		e.Val = float64(e.Val.(uint16))
	case int16:
		e.Val = float64(e.Val.(int16))
	case uint32:
		e.Val = float64(e.Val.(uint32))
	case int32:
		e.Val = float64(e.Val.(int32))
	case uint64:
		e.Val = float64(e.Val.(uint64))
	case int64:
		e.Val = float64(e.Val.(int64))
	case int:
		e.Val = float64(e.Val.(int))
	case float32:
		e.Val = float64(e.Val.(float32))
	case float64:
		e.Val = float64(e.Val.(float64))
	default:
		return errors.New("ArithmeticError: can only add numeric types")
	}
	return nil
}

func (e Element) Op(x Element, op string) (Element, error) {

	// Ensure e has float value
	err_e := e.AsFloat()
	if err_e != nil {
		panic(err_e)
	}

	// Ensure x has float value
	err_x := x.AsFloat()
	if err_x != nil {
		panic(err_x)
	}

	// Return no error
	return Ops[op](e.Val.(float64), x.Val.(float64)), nil
}

func (e Element) Add(x Element) (Element, error) {
	return e.Op(x, "+")
}

func (e Element) Prod(x Element) (Element, error) {
	return e.Op(x, "*")
}
