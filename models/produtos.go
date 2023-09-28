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
	query, err := db.Query("select * from produtos order by id asc")
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

func DelProduto(id string) {
	db := db.Connect()
	delProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	delProduto.Exec(id)
	defer db.Close()
}

func EditProduto(id string) Produto {
	db := db.Connect()
	produtoBanco, err := db.Query("select * from produtos where id = $1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Desc = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Qtd = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar

}

func UpdateProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.Connect()
	updateProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, qtd=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	updateProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
