package gocockroachdb_test

import (
	"testing"

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
			var c gocockroachdb.CockDB
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
