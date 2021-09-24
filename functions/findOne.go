package gohana

import (
	"fmt"
	"strings"
)

func (*Instance) FindOne(table string, columns []string, conditions map[string]string) (map[int]map[string]string, error) {
	results := make(map[int]map[string]string)

	var sql string

	if columns != nil {
		sqlColumns := strings.Join(columns, ", ")
		sql = fmt.Sprintf("SELECT %s FROM %s", sqlColumns, table)

	} else {
		sql = fmt.Sprintf("SELECT * FROM %s", table)
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

	sql = sql + " LIMIT 1"

	rows, err := Db.Query(sql)
	if err != nil {
		return nil, err
	}

	columnNames, _ := rows.Columns()
	columnCount := len(columnNames)
	values := make([]interface{}, columnCount)
	valuePtrs := make([]interface{}, columnCount)

	index := 0
	for rows.Next() {

		for i := range columnNames {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)
		results[index] = make(map[string]string)
		for i, col := range columnNames {
			val := values[i]

			b, ok := val.([]byte)
			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}

			results[index][col] = fmt.Sprintf("%v", v)
		}
		index++
	}

	return results, nil
}
