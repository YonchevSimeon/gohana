package gohana

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (*Instance) Max(table, column string, conditions map[string]string) (map[string]float64, error) {
	results := make(map[string]float64)

	var sql string

	if len(column) != 0 {
		sql = fmt.Sprintf("SELECT MAX(%s) FROM %s", column, table)
	} else {
		return results, errors.New("please specify a column")
	}

	if conditions != nil {
		index := 0
		for key, value := range conditions {
			if index == 0 {
				sql = sql + fmt.Sprintf(" WHERE %s = '%s'", key, value)
			} else {
				sql = sql + fmt.Sprintf(" AND %s = '%s'", key, value)
			}
			index++
		}
	}

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}

	columnNames, _ := rows.Columns()
	columnCount := len(columnNames)
	values := make([]interface{}, columnCount)
	valuePtrs := make([]interface{}, columnCount)

	for rows.Next() {

		for i := range columnNames {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		for i, col := range columnNames {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}

			valueString := fmt.Sprintf("%v", v)
			valueStringArgs := strings.Split(valueString, "/")
			firstValue, _ := strconv.ParseFloat(valueStringArgs[0], 64)
			var secondValue float64
			if len(valueStringArgs) > 1 {
				tempSecondValue, _ := strconv.ParseFloat(valueStringArgs[1], 64)
				secondValue = tempSecondValue
			}
			var dividedValue float64
			if secondValue != 0 {
				tempDividedValue := firstValue / secondValue
				dividedValue = tempDividedValue
			} else {
				dividedValue = firstValue
			}

			results[col] = dividedValue
		}
	}

	return results, nil
}
