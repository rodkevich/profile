//Module provides repository
//

package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx"

	"github.com/jackc/pgx/pgxpool"
	repo "github.com/rodkevich/profile/repository"
)

// Represents the repository model
type repository struct {
	db *pgxpool.Pool
}

// NewRepository will create a variable that represent the Repository struct
func NewRepository() (repo.Repository, error) {
	//get settings
	var config = os.Getenv("DATABASE_URL")
	//create data-base connection pool:
	pool, poolErr := pgxpool.Connect(context.Background(), config)
	if poolErr != nil {
		log.Fatalf("Unable to connection to database: %v\n", poolErr)
	}
	//defer pool.Close()
	log.Printf("Connected!")
	//config can be done here if from config
	//db.MaxIdleConns(idleConn)
	//db.MaxOpenConns(maxConn)
	return &repository{pool}, nil
}

// Up attaches the provider and create the person table
func (r *repository) Up() error {
	//create function in postgres: uuid_generate_v4() returns uuid
	ctx := context.Background()
	extension := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

	_, e := r.db.Query(ctx, extension)

	if e != nil {
		return e
	}
	// create table for persons
	query := `CREATE TABLE IF NOT EXISTS person (
              person_id uuid DEFAULT uuid_generate_v4 (),
              project_id VARCHAR NOT NULL,
              project_env VARCHAR NOT NULL,
              company_id VARCHAR,
              team_id VARCHAR NOT NULL,
              group_id VARCHAR NOT NULL,
              first_name VARCHAR NOT NULL,
              last_name VARCHAR NOT NULL,
              email VARCHAR UNIQUE NOT NULL,
              password VARCHAR,
              phone VARCHAR,
              description VARCHAR,
              PRIMARY KEY (person_id)
          )`

	_, eQ := r.db.Query(ctx, query)
	fmt.Println("Query was made")

	if eQ != nil {
		return eQ
	}
	return nil
}

// Close attaches the provider and close the connection
func (r *repository) Close() {
	r.db.Close()
}

// Create attaches the person repository and creating the data
func (r *repository) Create(person *repo.PersonModel) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var personId string
	query := "INSERT INTO person (" +
		"project_id, " +
		"project_env, " +
		"company_id, " +
		"team_id, " +
		"group_id, " +
		"first_name, " +
		"last_name, " +
		"email, " +
		"phone," +
		"password," +
		"description" +
		") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) " +
		"returning person_id;"
	err := r.db.QueryRow(
		ctx, query,
		person.ProjectID,
		person.ProjectEnv,
		person.CompanyID,
		person.TeamID,
		person.GroupID,
		person.FirstName,
		person.LastName,
		person.Email,
		person.Phone,
		person.Password,
		person.Description).Scan(&personId)
	if err != nil {
		return personId, err
	}
	fmt.Printf("Successfully created user with id %v\n", personId)
	return personId, nil
}

//Find attaches the person repository and finds all the data
func (r *repository) Find() ([]*repo.PersonModel, error) {
	persons := make([]*repo.PersonModel, 0)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "SELECT " +
		"person_id, " +
		"project_id, " +
		"project_env, " +
		"company_id, " +
		"team_id, " +
		"group_id, " +
		"first_name, " +
		"last_name, " +
		"email, " +
		"phone, " +
		"password, " +
		"description " +
		"FROM person"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		person := new(repo.PersonModel)
		err = rows.Scan(
			&person.ID,
			&person.ProjectID,
			&person.ProjectEnv,
			&person.CompanyID,
			&person.TeamID,
			&person.GroupID,
			&person.FirstName,
			&person.LastName,
			&person.Email,
			&person.Phone,
			&person.Password,
			&person.Description,
		)
		if err != nil {
			return nil, err
		}
		log.Println(person)
		persons = append(persons, person)
	}
	return persons, nil
}

// FindByID attaches the user repository and find data based on id
func (r *repository) FindByID(id string) (*repo.PersonModel, error) {
	person := new(repo.PersonModel)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := "SELECT person_id, first_name, last_name, email, phone FROM person WHERE person_id = $1"
	err := r.db.QueryRow(
		ctx,
		query,
		id).Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.Phone)
	if err != nil {
		return nil, err
	}
	return person, nil
}

// Update attaches the user repository and update data based on id
func (r *repository) Update(person *repo.PersonModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "UPDATE person SET " +
		"project_id = $1, " +
		"project_env = $2, " +
		"company_id = $3, " +
		"team_id = $4," +
		"group_id = $5," +
		"first_name = $6," +
		"last_name = $7, " +
		"email = $8, " +
		"phone = $9, " +
		"password = $10, " +
		"description = $11 " +
		"WHERE person_id = $12"

	err := r.db.QueryRow(
		ctx, query,
		person.ProjectID,
		person.ProjectEnv,
		person.CompanyID,
		person.TeamID,
		person.GroupID,
		person.FirstName,
		person.LastName,
		person.Email,
		person.Phone,
		person.Password,
		person.Description,
		person.ID).Scan()

	if err != nil {
		return err
	}
	return nil
}

//delete user table
func (r *repository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM person WHERE person_id=$1"
	qe := r.db.QueryRow(ctx, query, id)

	if qe != nil {
		log.Println(qe)
	}
	return nil
}

//truncate user table
func (r *repository) Truncate() error {
	ctx := context.Background()
	query := "TRUNCATE TABLE person;"
	stmt, err := r.db.Query(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

// Drop attaches the provider and drop the table
func (r *repository) Drop() error {
	var ctx = context.Background()
	var query string = "DROP TABLE IF EXISTS person"
	var stmt pgx.Rows
	var err error
	// query to db
	stmt, err = r.db.Query(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}
