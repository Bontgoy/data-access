package dbConnect

import (
	"database/sql"
	"fmt"

	"example.com/goDotEnvVariable"
	_ "github.com/lib/pq"
)

// var db *sql.DB

const (
	host   = "localhost"
	port   = 5432
	dbname = "learngo"
)

func ConnectionConfig() string {
	postgresqlDbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
		host, port, goDotEnvVariable.GetEnvVariable("DBUSER"), goDotEnvVariable.GetEnvVariable("DBPASSWORD"), goDotEnvVariable.GetEnvVariable("DATABASENAME"),
	)

	return postgresqlDbInfo
}

func TestConnection() *sql.DB {
	db, err := sql.Open("postgres", ConnectionConfig())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Established a successful connection!")
	return nil
}
