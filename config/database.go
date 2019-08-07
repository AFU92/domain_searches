package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const query string = `
	CREATE SEQUENCE IF NOT EXISTS domains_seq;

	CREATE SEQUENCE IF NOT EXISTS servers_seq;

	CREATE TABLE IF NOT EXISTS domains(
		id integer PRIMARY KEY DEFAULT nextval('domains_seq'),
		"name" varchar(250) UNIQUE NOT NULL,
		servers_changed BOOL DEFAULT FALSE,
		ssl_grade varchar(2),
		previous_ssl_grade varchar(2),
		logo varchar(300),
		title varchar(250),
		is_down BOOL DEFAULT FALSE,
		inserted_at TIMESTAMPTZ DEFAULT now()
	);

	CREATE TABLE IF NOT EXISTS servers(
		id integer PRIMARY KEY DEFAULT nextval('servers_seq'),
		"address" varchar(250) UNIQUE NOT NULL,
		ssl_grade varchar(2),
		country varchar(150),
		owner varchar(250),
		"current_status" BOOL DEFAULT TRUE,
		domain_id integer REFERENCES domains(id),
		inserted_at TIMESTAMPTZ DEFAULT now()
	);`

//ConnectToDB func that creates a DB connection
func ConnectToDB() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://ds_admin@localhost:26257/domain_search?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}

//CreatingDatabaseObject func that creates database objects if not exist
func CreateDatabaseObject() {
	db := ConnectToDB()
	defer db.Close()
	log.Println("Checking and creating necessary database objects")
	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("error creating database objects: ", err)
	}
}
