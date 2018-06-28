# SQL

## Цель


## План
* SELECT
* INSERT
* UPDATE
* DELETE
* CREATE TABLE
* JOIN, вложенные запросы
* Foreign Key
* Настройка psql

### Chapter 1

#### Table 1
* id SERIAL PRIMARY KEY
* title VARCHAR(64)
* author VARCHAR(64)

#### Table 2
* id SERIAL PRIMARY KEY
* first name VARCHAR(64)
* last name VARCHAR(64)
* email VARCHAR(64)

* id SERIAL PRIMARY KEY
* title VARCHAR(64)
* athor_id INTEGER REFERENCES authors(id)

### Chapter 1
```
> psql -U postgres
postgres-#

```


## Заключение


## Полезные ссылки
* https://eax.me/postgresql-install/
* https://4gophers.ru/articles/go-i-sql-bazy-dannyh/
* http://www.postgresqltutorial.com/postgresql-foreign-key/