package gohana

import "database/sql"

func (*Instance) Raw(query string) (*sql.Rows, error)  {
	
	return Db.Query(query)
	
}
