package db

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	//import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

var db *gorp.DbMap

//Init ...
func Init() {

	var err error
	db, err = ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB() (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", "postgres://tbzobuabpgyvwj:176b266a38ec5f4b4e84c09a2c81bdf3445e72de1a8dc314f161826385d4b070@ec2-54-75-249-162.eu-west-1.compute.amazonaws.com:5432/dc2iju1niatl10")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
