package frame

import (
	"fmt"
	"reflect"
)

type Frame struct {
	Columns map[string]Series
	Names   []string
}

func GoFrame(columns []interface{}) (Frame, error) {
	f := Frame{Columns: make(map[string]Series)}
	if columns == nil {
		return f, nil
	}

	f.Names = make([]string, len(columns))

	for i, c := range columns {
		switch c.(type) {
		case Series:
			f.Columns[c.(Series).Name] = c.(Series)
			f.Names[i] = c.(Series).Name
		case []interface{}:
			name := fmt.Sprintf("Column%d", i)
			data := make([]interface{}, len(c.([]interface{})))
			for i, col := range c.([]interface{}) {
				data[i] = col
			}
			series, err := GoSeries(name, data, reflect.TypeOf(data[0]))
			if err != nil {
				panic(err)
			}
			f.Columns[name] = series
			f.Names[i] = name
		default:
			fmt.Println(reflect.TypeOf(c))
			panic("Type not understood")
		}
	}
	return f, nil
}

func (frame *Frame) Get(column string) Series {
	return frame.Columns[column]
}
