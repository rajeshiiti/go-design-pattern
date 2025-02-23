package main

import "fmt"

type Database interface {
	SaveData(data string)
}

type NoSQLDatabase struct {
}

func (ns *NoSQLDatabase) InsertDocument(json string) {
	fmt.Println("No SQL database json- ", json)
}

type NoSQLAdapter struct {
	noSQLDatabase *NoSQLDatabase
}

func (nsa NoSQLAdapter) SaveData(data string) {
	fmt.Println("Adaptor converting request from string to nosql format")
	nsa.noSQLDatabase.InsertDocument(data)
}

func main() {
	noSQLDatabase := &NoSQLDatabase{}
	noSQLAdapter := &NoSQLAdapter{
		noSQLDatabase: noSQLDatabase,
	}
	noSQLAdapter.SaveData("Hello testibng data.")

}
