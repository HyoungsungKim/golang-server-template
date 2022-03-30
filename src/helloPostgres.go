package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func connectDB() {
	loadEnv()
	var (
		host     = os.Getenv("DB_ADDRES")
		port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DB")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	createStmt := `create table students(
		name varchar (20) Primary key,
		roll integer
	)`

	_, err = db.Exec(createStmt)

	selectStmt := `select * from Students`
	_, err = db.Exec(selectStmt)

	insertStmt := `insert into "students"("name", "roll") values('John', 1)`
	_, err = db.Exec(insertStmt)
	CheckError(err)

	insertDynStmt := `insert into "students"("name", "roll") values($1, $2)`
	_, err = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(err)
}

func loadEnv() {
	err := godotenv.Load("/app/.postgres_env")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
