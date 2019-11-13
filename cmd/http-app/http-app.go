package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	delivery "github.com/hobord/goddd1/delivery/http"
	persistence "github.com/hobord/goddd1/infrastructure/persistence/mysql"
	"github.com/hobord/goddd1/usecase"
)

func main() {
	r := mux.NewRouter()

	conn, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}

	repository := persistence.NewEntityMysqlRepository(conn)
	entityInteractor := usecase.NewExampleInteractor(&repository)

	delivery.MakeRouting(r, entityInteractor)

	http.ListenAndServe(":80", r)
}
