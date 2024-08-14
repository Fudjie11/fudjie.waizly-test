# sqldb

`sqldb` is a helper to connect to SQL database server.

It follows the design pattern where write operation will be directed to master DB while the read one to follower DB.
This pattern will improve the read scalability because follower DB can be scaled horizontally in easy way while still maintaining
data consistency which written to the master.

## Write to Master, Read from Follower

In `sqldb`, all `select`, `get` and `query` will come to `follower` database. All exec is coming to `master`. Or equivalent to `DDL` and `DML` is going to `master` and `data retrieval` is going to `follower`.

If follower DSN is not specified, then `follower` will be the same as `master`.

### Single Master - Follower model

`sqldb` only recognize one `master` database. This means `sqldb` is not suitable for `multi-master` database model.

### Usage example

For initializing
```go
dbclient, err := sqldb.Connect(context.Background(), sqldb.DBConfig{
    Driver:                "mysql", // depends on the engine, currently support postgresql and
    MasterDSN:             "username:password@tcp(host:port)/database_name",
    SlaveDSN:              "username:password@tcp(host:port)/database_name",
    MaxOpenConnections:    100, // if 0, it will use default config
    MaxIdleConnections:    10, // if 0, it will use default config
    ConnectionMaxLifetime: 10 * time.Second,
    Retry:                 3,
})
if err != nil {
    log.Println("error initialize db client, err : ", err)
    return
}
log.Println("database connected !!")
```

for any other example, you can check on example/sqlexample folder