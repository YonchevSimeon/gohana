package gohana

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func (*Instance) Insert(table string, intfc interface{}) (string, error) {

	res, err := json.MarshalIndent(&intfc, "", " ")
	if err != nil {
		return err.Error(), err
	}

	jsonString := string(res)
	pattern := `("[A-Za-z_]+"): ("*[A-Za-z0-9 ]+[.]*[0-9]*"*),*`
	pat := regexp.MustCompile(pattern)
	matches := pat.FindAllStringSubmatch(jsonString, -1)

	var columns []string
	var values string

	for _, match := range matches {
		columns = append(columns, match[1])

		if strings.HasPrefix(match[2], "\"") {
			values = values + strings.Replace(match[2], "\"", "'", 2) + ","
		} else {
			values = values + match[2] + ","
		}
	}
	values = values[:len(values)-1]
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)", table, strings.Join(columns, ", "), values)

	result, err := Db.Exec(sql)
	if err != nil {
		return err.Error(), err
	}

	var r, _ = result.RowsAffected()

	return fmt.Sprintf("Rows affected %d", r), nil
}
