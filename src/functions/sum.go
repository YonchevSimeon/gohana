package gohana

import (
	"errors"
	"fmt"
)

func (*Instance) Sum(table, column string, conditions map[string]string) (map[string]string, error) {
	results := make(map[string]string)

	var sql string

	if len(column) != 0 {
		sql = fmt.Sprintf("SELECT SUM(%s) FROM %s", column, table)
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
			results[col] = valueString
		}
	}

	return results, nil
}
