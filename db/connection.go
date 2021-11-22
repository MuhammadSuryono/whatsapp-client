package db

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Connection *gorm.DB

type ConnectionHandler struct {
	DB_HOST   string
	DB_PORT   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	TIMEZONE  string
	DB_DRIVER string
}

func (con ConnectionHandler) CreateNewConnection() {
	switch con.DB_DRIVER {
	case "mysql":
		con.mysqlConnection()
		break
	case "postgres":
		con.postgresConnection()
		break
	case "sql-server":
		con.sqlServerConnection()
		break
	default:
		log.Fatalf("Driver database not set")
	}
}

func CloseConnectionDb(conn *gorm.DB) {
	connDb, err := conn.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed close database: %v", err))
	}

	connDb.Close()
}

// TODO: Init koneksi ini di root file ex. main.go
func InitConnectionFromEnvirontment() ConnectionHandler {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	return ConnectionHandler{
		DB_HOST:   host,
		DB_PORT:   port,
		DB_USER:   user,
		DB_PASS:   pass,
		DB_NAME:   dbName,
		DB_DRIVER: dbDriver,
	}
}

func InitConnection(host, port, user, pass, dbName, dbDriver string) ConnectionHandler {
	return ConnectionHandler{
		DB_HOST:   host,
		DB_PORT:   port,
		DB_USER:   user,
		DB_PASS:   pass,
		DB_NAME:   dbName,
		DB_DRIVER: dbDriver,
	}
}

func (con ConnectionHandler) mysqlConnection() {
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		con.DB_USER, con.DB_PASS, con.DB_HOST, con.DB_PORT, con.DB_NAME)
	db, err := gorm.Open(mysql.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}

func (con ConnectionHandler) postgresConnection() {
	args := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		con.DB_HOST, con.DB_USER, con.DB_PASS, con.DB_NAME, con.DB_PORT)
	db, err := gorm.Open(postgres.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}

func (con ConnectionHandler) sqlServerConnection() {
	args := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		con.DB_USER, con.DB_PASS, con.DB_HOST, con.DB_PORT, con.DB_NAME)
	db, err := gorm.Open(sqlserver.Open(args))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect database with setting: %s", args))
	}
	Connection = db

	log.Info(fmt.Sprintf("Database %s Connected", con.DB_DRIVER))
}
