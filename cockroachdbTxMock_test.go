package gocockroachdb_test

import (
	"testing"

	gocockroachdb "github.com/GolangToolKits/go-cockroachdb"
)

func TestCockDbTxMock_Insert(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
		want2 int64
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
			var c gocockroachdb.CockDbMock
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
			c.MockInsertID1 = 1
			c.MockInsertSuccess1 = true

			db := c.New()
			//gotc := db.Connect()
			ctx := db.BeginTransaction()
			got, got2 := ctx.Insert(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			if got2 <= 0 {
				t.Errorf("Insert() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCockDbTxMock_Update(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
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
			var c gocockroachdb.CockDbMock
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"

			c.MockUpdateSuccess1 = true

			db := c.New()
			//gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Update(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbTxMock_Delete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
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
			var c gocockroachdb.CockDbMock
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"

			c.MockDeleteSuccess1 = true

			db := c.New()
			//gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Delete(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbTxMock_Commit(t *testing.T) {
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
			var c gocockroachdb.CockDbMock
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"

			c.MockCommitSuccess = true

			db := c.New()
			//gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Commit()
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Commit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbTxMock_Rollback(t *testing.T) {
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
			var c gocockroachdb.CockDbMock
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"

			c.MockRollbackSuccess = true

			db := c.New()
			//gotc := db.Connect()
			ctx := db.BeginTransaction()
			got := ctx.Rollback()
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Rollback() = %v, want %v", got, tt.want)
			}
		})
	}
}
