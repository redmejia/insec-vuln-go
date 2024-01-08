# insec-vuln-go


Requirement 
- Go(Golang) install instructin -> https://go.dev/doc/install
- Docker Install documentation -> https://docs.docker.com/engine/install/
- Reactjs -> https://react.dev


Build and Run
1. Build and run container 
```bash
    make build_up
```
Restart 
2. Restart continers

```bash
    make restart
```

Database\
Postgres database go to user-service/internal/migration/create.sql.
To persist database on your local machine create a directory on your root application
```bash
    mkdir -p vuln-go/data 
```


For persist db create a directory 
