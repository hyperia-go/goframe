package frame

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadCsv(file string) (Frame, error) {
	empty, _ := GoFrame(make([]interface{}, 0))
	f, err := os.Open(file)
	if err != nil {
		return empty, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	row, err := csvr.Read()
	columns := make([][]interface{}, len(row))
	if err != nil {
		if err == io.EOF {
			err = nil
			return empty, nil
		}
		return empty, err
	}

	for i, v := range row {
		columns[i] = make([]interface{}, 1)
		columns[i][0] = v
	}

	// Read until EOF
	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return empty, err
			}
		}

		for i, v := range row {
			columns[i] = append(columns[i], v)
		}
	}

	series_columns := make([]Series, len(columns))
	for i, col := range columns {
		name := col[0]
		data := col[1:]
		s, err := GoSeries(name.(string), data)
		if err != nil {
			return empty, err
		}
		series_columns[i] = s
	}
	return GoFrameFromSeries(series_columns)
}
