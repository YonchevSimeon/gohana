package gohana

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func (*Instance) Update(table string, conditions map[string]string, intfc interface{}) (string, error) {
	sql := fmt.Sprintf("UPDATE %s", table)

	res, err := json.MarshalIndent(&intfc, "", " ")
	if err != nil {

		fmt.Println(err)
	}

	jsonString := string(res)
	pattern := `"([A-Za-z_]+)": ("*[A-Za-z0-9 ]+[.]*[0-9]*"*),*`
	pat := regexp.MustCompile(pattern)
	matches := pat.FindAllStringSubmatch(jsonString, -1)

	index := 0
	var v string
	for _, match := range matches {

		if strings.HasPrefix(match[2], "\"") {
			v = strings.Replace(match[2], "\"", "", 2)
		} else {
			v = match[2]
		}

		if index == 0 {
			sql = sql + fmt.Sprintf(" SET %s = '%s'", match[1], v)
		} else {
			sql = sql + fmt.Sprintf(", %s = '%s'", match[1], v)
		}
		index++
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

	result, err := Db.Exec(sql)
	if err != nil {
		return err.Error(), err
	}

	var r, _ = result.RowsAffected()

	return fmt.Sprintf("Rows affected %d", r), nil
}
