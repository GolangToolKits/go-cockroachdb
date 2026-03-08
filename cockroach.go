package gocockroachdb

import (
	"context"

	// CockroachDB helper library
	// "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5" // CockroachDB helper library

	// crdb "github.com/cockroachdb/cockroach-go/v2/crdb"
	"log"
	// PostgreSQL driver

	"github.com/jackc/pgx/v5/pgxpool"
)

// MyDB MyDB
type CockDB struct {
	Host        string
	User        string
	Password    string
	Database    string
	Port        string
	Sslmode     string
	Sslrootcert string
	//db          *sql.DB
	err    error
	dbpool *pgxpool.Pool
	//crdb crdb
	//err error
}

// Connect Connect
func (c *CockDB) Connect() bool {
	var rtn = false
	ctx := context.Background()
	var conStr string
	//postgresql://{username}:{password}@{host}:{port}/{database}?sslmode=verify-full&sslrootcert={root-cert}

	//postgresql://root@localhost:26257/defaultdb?sslmode=disable
	if c.Host == "localhost" {
		conStr = "postgresql://" + c.User + ":" + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=" + c.Sslmode
	} else {
		conStr = "postgresql://" + c.User + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Database + "?sslmode=" + c.Sslmode + "&sslrootcert=" + c.Sslrootcert
	}

	// conStr = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"

	c.dbpool, c.err = pgxpool.New(ctx, conStr)
	if c.err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", c.err)
	}
	// defer dbpool.Close() // Ensure the pool is closed when the application shuts down

	// Verify the connection
	err2 := c.dbpool.Ping(ctx)
	if err2 != nil {
		log.Fatalf("Ping failed: %v\n", err2)
	} else {
		rtn = true
	}

	// var sql = "insert into orders (order_number, entered) values($1, $2) RETURNING id"

	// var lastInsertID int64
	// // _, err = conn.Exec(ctx, "insertOrder", "123_GGQ", time.Now()).Scan(&lastInsertID)
	// err3 := c.dbpool.QueryRow(ctx, sql, "123_GGQ", time.Now()).Scan(&lastInsertID)
	// if err3 != nil {
	// 	log.Fatalf("Unable to execute statement: %v", err3)
	// }

	// if lastInsertID > 0 {
	// 	// id = lastInsertID
	// 	rtn = true
	// }

	return rtn
}

// Insert Insert requires use of  RETURNING id on end of insert query
func (c *CockDB) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	ctx := context.Background()

	err3 := c.dbpool.QueryRow(ctx, query, args...).Scan(&id)
	if err3 != nil {
		log.Fatalf("Unable to execute statement: %v", err3)
	}

	if id > 0 {
		// id = lastInsertID
		success = true
	}
	// var stmtIns *sql.Stmt
	// stmtIns, p.err = p.db.Prepare(query)
	// if p.err != nil {
	// 	log.Println("Error:", p.err.Error())
	// } else {
	// 	defer stmtIns.Close()
	// 	var lastInsertID int64
	// 	err := stmtIns.QueryRow(args...).Scan(&lastInsertID)
	// 	if err != nil {
	// 		log.Println("Insert Exec err:", err.Error())
	// 	} else {
	// 		if lastInsertID > 0 {
	// 			id = lastInsertID
	// 			success = true
	// 		}
	// 		// success = true
	// 	}
	// }
	return success, id
}

func (c *CockDB) Close() bool {
	c.dbpool.Close()
	return true
}

func (c *CockDB) New() Database {
	return c
}
