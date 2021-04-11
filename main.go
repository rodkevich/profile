package main

import (
	"fmt"
	"log"
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

	//persons list
	var persons []*repository.PersonModel

	//some test persons
	var Peter = &repository.PersonModel{
		ProjectID:   "S0me_pr0j3ct_1d",
		ProjectEnv:  "S0me_pr0j3ct_3nv",
		CompanyID:   "S0me_c0mp4ny_1d",
		TeamID:      "S0me_t34m_1d",
		GroupID:     "S0me_gr0up_1d",
		FirstName:   "Peter",
		LastName:    "Bar",
		Email:       "Peter@bar.foo",
		Phone:       "08123456789",
		Description: "Some Description 1",
	}

	var John = &repository.PersonModel{
		ProjectID:   "S0me_pr0j3ct_1d",
		ProjectEnv:  "S0me_pr0j3ct_3nv",
		CompanyID:   "S0me_c0mp4ny_1d",
		TeamID:      "S0me_t34m_1d",
		GroupID:     "S0me_gr0up_2d",
		FirstName:   "John",
		LastName:    "Foo",
		Email:       "John@foo.bar",
		Phone:       "9887654210",
		Description: "Some Description 2",
	}

	//append to list
	persons = append(persons, Peter)
	persons = append(persons, John)

	for _, person := range persons {
		id, _ := rep.Create(person)
		log.Println(id)
	}
	users, err := rep.Find()
	if err != nil {
		log.Println(err)
	}
	log.Println(users)

	var JohnUpdated = &repository.PersonModel{
		ID:          users[0].ID,
		ProjectID:   "S0me_pr0j3ct_1d",
		ProjectEnv:  "S0me_pr0j3ct_3nv",
		CompanyID:   "S0me_c0mp4ny_1d",
		TeamID:      "S0me_t34m_1d",
		GroupID:     "S0me_gr0up_1d",
		FirstName:   "UpdatedJohn",
		LastName:    "UpdatedDooooo",
		Email:       "UpdatedJohn@Dooooo.Bar",
		Phone:       "Updated9887654210",
		Password:    "UpdatedPass_01-1",
		Description: "V3ry_n1c3_u53r",
	}

	_ = rep.Update(JohnUpdated)

	var J *repository.PersonModel

	J, err = rep.FindByID(users[0].ID)
	if err != nil {
		log.Panic(err)
	}

	log.Println(J)
	//
	//for i := 1; i <= 10; i++ {
	//	_, _ = rep.Create(John)
	//}

	////delete user by id
	//var u = users[1].ID
	//err = rep.Delete(u)

	//// truncate the whole table
	//trResp := rep.Truncate()
	//fmt.Println("trResp", trResp)

	////delete the whole table
	//_ = rep.Drop()

	fmt.Println("Done")
}
