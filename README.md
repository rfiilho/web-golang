# Web com GoLang

Aplicação web desenvolvida com GoLang.

## Features

- CRUD
- Padrão MVC

## Requisitos

- Go
- PostgreSQL

## Instalação

Criação da tabela "produtos"

```sh
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	qtd integer
)
```

Editar as credenciais de conexão com o banco de dados.

```sh
vi db/db.go
```