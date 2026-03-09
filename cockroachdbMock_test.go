package gocockroachdb_test

import (
	"testing"

	gocockroachdb "github.com/GolangToolKits/go-cockroachdb"
)

func TestCockDbMock_Connect(t *testing.T) {
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

			c.MockConnectSuccess = tt.want
			db := c.New()
			gotc := db.Connect()
			// TODO: update the condition below to compare got with tt.want.
			if !gotc {
				t.Errorf("Connect() = %v, want %v", gotc, tt.want)
			}
		})
	}
}

func TestCockDbMock_New(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want gocockroachdb.Database
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			//want: true,
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
			got := c.New()
			// TODO: update the condition below to compare got with tt.want.
			if got == nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_BeginTransaction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want gocockroachdb.Transaction
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			//want: true,
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
			db := c.New()
			got := db.BeginTransaction()
			// TODO: update the condition below to compare got with tt.want.
			if got == nil {
				t.Errorf("BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_Test(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  *gocockroachdb.DbRow
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			//want: true,
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
			c.MockTestRow = &gocockroachdb.DbRow{}
			db := c.New()
			got := db.Test(tt.query, tt.args...)
			// TODO: update the condition below to compare got with tt.want.
			if got == nil {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_Insert(t *testing.T) {
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
			//want: true,
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
			got, got2 := db.Insert(tt.query, tt.args...)
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
			if got2 != 1 {
				t.Errorf("Insert() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestCockDbMock_Update(t *testing.T) {
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
			//want: true,
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
			got := db.Update(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if !got {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_Get(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  *gocockroachdb.DbRow
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDbMock
			got := c.Get(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_GetList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  *gocockroachdb.DbRows
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDbMock
			got := c.GetList(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_Delete(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		query string
		args  []any
		want  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDbMock
			got := c.Delete(tt.query, tt.args)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCockDbMock_Close(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var c gocockroachdb.CockDbMock
			got := c.Close()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
