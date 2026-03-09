# go-cockroachdb
A simple mockable postgres db interface


[![Go Report Card](https://goreportcard.com/badge/github.com/GolangToolKits/go-cockroachdb)](https://goreportcard.com/report/github.com/GolangToolKits/go-cockroachdb)

```go

    // Connecting to local insecure CockroachDB
    // For production cluster add password and set SslMode and Sslrootrert
     cc := &CockDB{ 
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
    }

  
    c := cc.New()
    c.Connect()
    //Insert requires  RETURNING id on end of prepared statement query
        //Example: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id"
    c.Insert(query, args...)
    c.Update(query, args...)
    c.Get(query, args...)
    c.GetList(query, args...)
    c.Delete(query, args...)

```

## Also Supports Transactions

```go

    cc := &CockDB{ 
			c.Database = "customer_orders"
			c.Host = "localhost"
			c.Port = "26257"
			c.User = "root"
			c.Sslmode = "disable"
    }

  
    c := cc.New()
    c.Connect()
    tr := c.BeginTransaction()
    //Insert requires  RETURNING id on end of prepared statement query
        //Example: "insert into carrier (carrier, type, store_id) values($1, $2, $3) RETURNING id"
    tr.Insert(query, args...)
    tr.Update(query, args...)
    tr.Get(query, args...)
    tr.GetList(query, args...)
    tr.Delete(query, args...)
    tr.Commit()
    //Or
    tr.Rollback()


```


