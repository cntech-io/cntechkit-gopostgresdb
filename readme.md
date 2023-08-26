### postgresdb helper methods for go projects

```go
postgresdb := NewPostgresDB().
    Connect()

postgresdb.Do().FindOne()

postgresdb.Migrate(...entities)
```
