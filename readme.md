### install

```bash
go get github.com/cntech-io/cntechkit-gopostgresdb/v2
```

### Methods

| Method                                               | Description                                       |
| ---------------------------------------------------- | ------------------------------------------------- |
| postgresdb.NewPostgresDB()                           | Creates postgresdb instance                       |
| &nbsp;&nbsp;&nbsp;&nbsp;.Connect()                   | Connects to mongodb                               |
| &nbsp;&nbsp;&nbsp;&nbsp;.Migrate(dst ...interface{}) | Migrate entities (see gorm)                       |
| &nbsp;&nbsp;&nbsp;&nbsp;.Do()                        | Enables to react mongodb methods                  |
| env.NewPostgresDBEnv()                               | Loads predefined postgresdb environment variables |
