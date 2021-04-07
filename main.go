package main

import (
	"os"

	"github.com/rodkevich/profile/repository"
	"github.com/rodkevich/profile/repository/postgres"
)

func main() {
	_ = os.Setenv("DATABASE_URL", "postgresql://postgres:postgres@0.0.0.0:5432/postgres")

	var (
		rep repository.Repository
		err error
	)

	rep, err = postgres.NewRepository() //init postgres rep db

	if err != nil {
		panic(err)
	}

	//create tables if not presented
	err = rep.Up()
	if err != nil {
		panic(err)
	}
}
