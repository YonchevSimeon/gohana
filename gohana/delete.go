package gohana

import (
	"errors"
	"fmt"
)

func (*Instance) Delete(table string, conditions map[string]string) (string, error) {

	sql := fmt.Sprintf("DELETE FROM %s", table)

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
	} else {
		return "please specify conditions", errors.New("please specify conditions")
	}

	result, err := Db.Exec(sql)
	if err != nil {
		return err.Error(), err
	}

	var r, _ = result.RowsAffected()

	return fmt.Sprintf("Rows affected %d", r), nil
}
