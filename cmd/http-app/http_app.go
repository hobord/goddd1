package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	httpdelivery "github.com/hobord/goddd1/delivery/http"
	persistence "github.com/hobord/goddd1/infrastructure/persistence/mysql"
	"github.com/hobord/goddd1/usecase"
)

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	dbConnection := os.Getenv("DB_CONNECTION")
	if dbConnection == "" {
		dbConnection = "user:password@/dbname"
	}

	conn, err := sql.Open("mysql", dbConnection)
	if err != nil {
		log.Fatal(err)
	}

	repository := persistence.NewEntityMysqlRepository(conn)
	interactor := usecase.NewExampleInteractor(repository)

	r := mux.NewRouter()

	httpdelivery.MakeRouting(r, interactor)

	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
