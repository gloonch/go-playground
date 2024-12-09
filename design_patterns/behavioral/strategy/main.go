package main

import "fmt"

type IDBConnection interface {
	Connect()
}

type DBConnection struct {
	Db IDBConnection
}

func (con DBConnection) DBConnect() {
	con.Db.Connect()
}

type MySqlConnection struct {
	ConnectionString string
}

func (c MySqlConnection) Connect() {
	fmt.Println("MySqlConnection " + c.ConnectionString)
}

type PostgresConnection struct {
	ConnectionString string
}

func (c PostgresConnection) Connect() {
	fmt.Println("PostgresConnection " + c.ConnectionString)
}

type MongoDBConnection struct {
	ConnectionString string
}

func (c MongoDBConnection) Connect() {
	fmt.Println("MongoDBConnection " + c.ConnectionString)
}

func main() {
	mySqlConnection := MySqlConnection{ConnectionString: "MySQL DB is connected"}
	conn := DBConnection{Db: mySqlConnection}
	conn.DBConnect()

	pgConnection := PostgresConnection{ConnectionString: "Postgres DB is connected"}
	connPg := DBConnection{Db: pgConnection}
	connPg.DBConnect()

	mongoConnection := MongoDBConnection{ConnectionString: "MongoDB is connected"}
	connMongo := DBConnection{Db: mongoConnection}
	connMongo.DBConnect()

}
