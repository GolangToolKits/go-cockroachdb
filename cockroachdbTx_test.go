package gocockroachdb_test

import (
	"testing"
	"time"

	gocockroachdb "github.com/GolangToolKits/go-cockroachdb"
)

func TestCockDbTx_Insert(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query    string
		args     []any
		want     bool
		want2    int64
		rollback bool
	}{
		// TODO: Add test cases.
		{
			name:  "test 1",
			query: "insert into orders (order_number, entered) values($1, $2) RETURNING id",

			want:  true,
			want2: 0,
			args:  []any{"123_GGQl_WWER3T", time.Now()},
		},
		{
			name:  "test 2",
			query: "insert into orders (order_number, entered) values($1, $2) RETURNING id",

			want:     true,
			want2:    0,
			args:     []any{"123_GGQl_WWER3T_ttt", time.Now()},
			rollback: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.

			var c gocockroachdb.CockDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			ctx := db.BeginTransaction()

			//var ct gocockroachdb.CockDbTx
			got, got2 := ctx.Insert(tt.query, tt.args...)

			if tt.rollback {
				ctx.Rollback()
			} else {
				ctx.Commit()
			}

			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			if got != true {
				t.Errorf("Insert() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCockDbTx_Update(t *testing.T) {
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
			args:  []any{"123_GGQl_WWERTT----3", time.Now(), 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Update(tt.query, tt.args...)

			ctx.Commit()

			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbTx_Delete(t *testing.T) {
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
			args:  []any{30},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDB
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			db := c.New()
			gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Delete(tt.query, tt.args...)
			ctx.Commit()

			db.Close()
			// TODO: update the condition below to compare got with tt.want.
			if gotc != true || got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
