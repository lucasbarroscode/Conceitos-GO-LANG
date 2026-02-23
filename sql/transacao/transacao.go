package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//o Begin inicia a trasacao
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	stmt.Exec(4000, "Bia")
	stmt.Exec(4001, "Carlos")
	//como eh uma transao, se fosse fora da transacao a bia e o carlos seriam inseridos, mas como Tiago deu errado, não é inserido nenhum
	//pq eh tratado logo abaixo
	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
