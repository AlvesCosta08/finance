package main


import (
	"database/sql"
	"log"
	"github.com/AlvesCosta08/finance/api"
	db "github.com/AlvesCosta08/finance/db/sqlc"
	_ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main()  {
	conn,err := sql.Open(dbDriver,dbSource)
	if err != nil {
		log.Fatal("cannot connect to db : ", err)
	}
	
	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api : ", err)
	}
}