## Useful commands

````bash
go env | grep PATH # show go path
go mod init github.com/tupizz/curso-go ## inicia modulo
go mod tidy ## instala dependencias e requirements, remove pacotes que não estamos utilizandos
go mod tidy -e ## install all dependencies except those that are not used and found in the go.mod file
go get golang.org/x/exp/constraints ## instala determinada biblioteca

# workspaces
go mod edit -replace=github.com/tupizz/curso-go=../curso-go ## substitui o caminho do pacote
go work init ./math ./sistema ## inicia uma nova workspace com os pacotes math e sistema que enxergam um ao outro quando trabalhamos com pacotes locais

# running tests
go test ./... ## run all tests
go test -v ./... ## run all tests with verbose
go test -coverprofile=coverage.out # run all tests and generate coverage file
go tool cover --html=coverage.out # open coverage file in browser to check what is missing
go test -bench=. # run benchmarks in all tests benchmarking created
go test -bench=. -run=^# # run benchmarks in all tests benchmarking created and run only benchmarks that starts with 
go test -bench=. -run=^# -benchmem # run benchmarks in all tests benchmarking created and run only benchmarks that starts with and show memory allocation
go help test # show all test documentation
go test -fuzz=. -fuzztime=10s # run fuzzing tests for 10 seconds
````

### mysql cheat sheet

```bash
# bash 
docker-compose exec mysql bash
mysql -uroot -p goexpert ## senha: root

# sql
drop database goexpert;
create database goexpert;
create table products (id varchar(255), name varchar(255), price decimal(10,2), primary key (id))
```

## Documentations

### Templates
https://pkg.go.dev/html/template