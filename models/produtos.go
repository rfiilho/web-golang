package models

import (
	"web/db"
)

type Produto struct {
	Id    int
	Nome  string
	Desc  string
	Preco float64
	Qtd   int
}

func GetProdutos() []Produto {
	db := db.Connect()
	query, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, qtd int
		var nome, desc string
		var preco float64

		err = query.Scan(&id, &nome, &desc, &preco, &qtd)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Id = id
		p.Desc = desc
		p.Qtd = qtd
		p.Preco = preco

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func NovoProduto(nome, descricao string, preco float64, quantidade int) []Produto {
	db := db.Connect()
	insertDados, err := db.Prepare("insert into produtos(nome, descricao, preco, qtd) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
	return nil
}
