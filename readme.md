## Useful commands

````bash
go mod init github.com/tupizz/curso-go ## inicia modulo
go mod tidy ## instala dependencias e requirements, remove pacotes que não estamos utilizandos
go get golang.org/x/exp/constraints ## instala determinada biblioteca
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