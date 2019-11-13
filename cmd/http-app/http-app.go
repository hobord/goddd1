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

	r := mux.NewRouter()

	conn, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal("Cant connect to database: %v", err)
	}

	repository := persistence.NewEntityMysqlRepository(conn)
	entityInteractor := usecase.NewExampleInteractor(&repository)

	httpdelivery.MakeRouting(r, entityInteractor)

	log.Fatal(http.ListenAndServe(":"+httpPort, r))
}
