package db

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pressly/goose"
)

//Server ...
type Server struct {
	Env string
}

//factory
var (
	DBConn *sql.DB
	flags  = flag.NewFlagSet("goose", flag.ExitOnError)
	dir    = flags.String("dir", "./migrations/", "directory with migration files")
)

//InitDb ...
func InitDb() {

	a := Server{}
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var (
		err            error
		maxLifeTimeInt int
		maxIdleConns   int
		maxOpenConns   int
	)
	maxLifeTimeInt, _ = strconv.Atoi(os.Getenv("CONNMAXLIFETIME"))
	maxIdleConns, _ = strconv.Atoi(os.Getenv("MAXIDLECONNS"))
	maxOpenConns, _ = strconv.Atoi(os.Getenv("MAXOPENCONNS"))

	maxLifeTime := time.Duration(maxLifeTimeInt)

	DBConn, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("[db/init] - Error when trying to open connection (%s). Error: %s", a.Env, err.Error())
	}
	DBConn.SetConnMaxLifetime(time.Minute * maxLifeTime)
	DBConn.SetMaxIdleConns(maxIdleConns)
	DBConn.SetMaxOpenConns(maxOpenConns)

	goose.SetDialect("mysql")
	if err := goose.Up(DBConn, "./migrations"); err != nil {
		log.Println("[db/init] - goose", err)
	}
}

//GetDNS representa a recuperação do acesso ao banco
func (a *Server) GetDNS() string {
	var (
		dbUser     string
		dbPassword string
		dbname     string
		dbHost     string
		dbPort     int
	)

	dbUser = os.Getenv("DBUSER")
	dbPassword = os.Getenv("DBPASSWORD")
	dbname = os.Getenv("DBNAME")
	dbHost = os.Getenv("DBHOST")
	dbPort, _ = strconv.Atoi(os.Getenv("DBPORT"))

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbname)
	return connectionString
}
