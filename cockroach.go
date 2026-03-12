package gocockroachdb

import (
	"database/sql"

	"log"
	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// CockDB CockDB
type CockDB struct {
	Host        string
	User        string
	Password    string
	Database    string
	Port        string
	Sslmode     string
	Sslrootcert string
	db          *sql.DB
	err         error
}

// Connect Connect
func (c *CockDB) Connect() bool {
	var rtn = false
	var conStr string
	//postgresql://{username}:{password}@{host}:{port}/{database}?sslmode=verify-full&sslrootcert={root-cert}

	//postgresql://root@localhost:26257/defaultdb?sslmode=disable
	if c.Host == "localhost" {
		conStr = "postgresql://" + c.User + ":" + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=" + c.Sslmode
	} else {
		conStr = "postgresql://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=" + c.Sslmode + "&sslrootcert=" + c.Sslrootcert
	}
	log.Println("constr: ", conStr)

	c.db, c.err = sql.Open("postgres", conStr)
	if c.err == nil {
		c.err = c.db.Ping()
		if c.err != nil {
			log.Println("Database Connect Error:", c.err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}

// Test Test
func (c *CockDB) Test(query string, args ...any) *DbRow {
	return c.Get(query, args...)
}

// BeginTransaction BeginTransaction
func (c *CockDB) BeginTransaction() Transaction {
	var trans Transaction
	var pgTrans CockDbTx
	tx, err := c.db.Begin()
	if err == nil && tx != nil {
		pgTrans.Tx = tx
	}
	trans = &pgTrans
	return trans
}

// Insert Insert requires use of  RETURNING id on end of insert query
func (c *CockDB) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	stmtIns, err := c.db.Prepare(query)
	if err != nil {
		log.Println("Error:", c.err.Error())
	} else {
		defer stmtIns.Close()
		var lastInsertID int64
		err := stmtIns.QueryRow(args...).Scan(&lastInsertID)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			if lastInsertID > 0 {
				id = lastInsertID
				success = true
			}
		}
	}
	return success, id
}

// Update Update
func (c *CockDB) Update(query string, args ...any) bool {
	var success = false
	stmtUp, err := c.db.Prepare(query)
	if err != nil {
		log.Println("Error:", c.err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records updated")
			} else {
				success = true
			}
		}
	}
	return success
}

// Get Get
func (c *CockDB) Get(query string, args ...any) *DbRow {
	var rtn DbRow
	stmtGet, err := c.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		if err != nil {
			log.Println("Get err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rtn.Row = append(rtn.Row, value)
					}
				}
			}
		}
	}
	return &rtn
}

// GetList GetList
func (c *CockDB) GetList(query string, args ...any) *DbRows {
	var rtn DbRows
	stmtGet, err := c.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		if err != nil {
			log.Println("GetList err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					var rowValuesStr []string
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rowValuesStr = append(rowValuesStr, value)
					}
					rtn.Rows = append(rtn.Rows, rowValuesStr)
				}
			}
		}
	}
	return &rtn
}

// Delete Delete
func (c *CockDB) Delete(query string, args ...any) bool {
	var success = false
	stmt, err := c.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records deleted")
			} else {
				success = true
			}
		}
	}
	return success
}

func (c *CockDB) Exec(query string) bool {
	var rtn bool
	var a []any
	rs, err := c.db.Exec(query, a...)
	if err == nil {
		log.Println(rs)
		rtn = true
	}
	return rtn
}

func (c *CockDB) Close() bool {
	var rtn = false
	err := c.db.Close()
	if err == nil {
		rtn = true
	}
	return rtn
}

func (c *CockDB) New() Database {
	return c
}
