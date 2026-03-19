package gocockroachdb_test

import (
	"testing"
	"time"

	gocockroachdb "github.com/GolangToolKits/go-cockroachdb"
)

func TestCockDB_Connect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			got := db.Connect()
			// TODO: update the condition below to compare got with tt.want.
			if got == false {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
			db.Close()
		})
	}
}

func TestCockDB_Test(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  int //*gocockroachdb.DbRow
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "select count(*) from orders",
			//args:  []any{3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := db.Test(tt.query, tt.args...)
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || len(got.Row) != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_Insert(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		arg1  string
		arg2  time.Time
		want  bool
		want2 int64
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "insert into orders (order_number, entered) values($1, $2) RETURNING id",
			arg1:  "123_GGQl_WWER2",
			arg2:  time.Now(),
			want:  true,
			want2: 0,
			args:  []any{"123_GGQl_WWER1", time.Now()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			//var a []any
			//a = append(a, tt.arg1, tt.arg2)
			//tt.args = a
			db := c.New()
			gotc := db.Connect()
			got, got2 := db.Insert(tt.query, tt.args...)
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true {
				t.Errorf("Connect failed")
			}
			if got != true {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			if got2 <= 0 {
				t.Errorf("Insert() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCockDB_Update(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "update orders set order_number = $1 , updated = $2 where id = $3 ",
			args:  []any{"123_GGQl_WWER----3", time.Now(), 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := db.Update(tt.query, tt.args...)
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || got != true {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_Get(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		//want  *gocockroachdb.DbRow
		want int
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "select id, order_number, entered, updated from orders where id = $1",
			args:  []any{3},
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := db.Get(tt.query, tt.args...)
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || len(got.Row) != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_GetList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  int //*gocockroachdb.DbRows
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "select * from orders where order_number = $1",
			args:  []any{"123_GGQl_WWER"},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := db.GetList(tt.query, tt.args...)
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || len(got.Rows) != tt.want {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_Delete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "delete from orders where id = $1",
			args:  []any{13},
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := c.Delete(tt.query, tt.args...)
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || got != false {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_BeginTransaction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want gocockroachdb.Transaction
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got := db.BeginTransaction()
			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || got == nil {
				t.Errorf("BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDB_Exec(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query1 string
		query2 string
		query3 string
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			query1: "create database IF NOT EXISTS roachtest6; ",
			query2: "use roachtest6",
			query3: "CREATE TABLE IF NOT EXISTS public.orders (" +
				"id INT8 NOT NULL GENERATED ALWAYS AS IDENTITY," +
				"order_number VARCHAR(20) NOT NULL," +
				"entered DATE NOT NULL," +
				"updated DATE NULL," +
				"CONSTRAINT orders_pkey PRIMARY KEY (id ASC)" +
				" ) WITH (schema_locked = true);",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			var c gocockroachdb.CockRDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			got, _ := db.Exec(tt.query1)
			got2, _ := db.Exec(tt.query2)
			got3, _ := db.Exec(tt.query3)

			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if !gotc || !got || !got2 || !got3 {
				t.Errorf("CreateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
