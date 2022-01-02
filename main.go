package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/iramarfalcao/codebank/infrastructure/repository"
	"github.com/iramarfalcao/codebank/usecase"
)

func main() {
	db := setupDb()
	defer db.Close()
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	// useCase.KafkaProducer = producer
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("user"),
		os.Getenv("password"),
		os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
