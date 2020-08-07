package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Profile struct {
	ProfileId int
	FirstName string
	LastName  string
	Age       int
}

func main() {
	writeDB()
	readDB()

}

func writeDB() {
	db, err := sql.Open("sqlite3", "/tmp/personal.db")
	checkError(err)
	statement, err := db.Prepare("insert into Profile (FirstName, LastName, Age) values(?, ?, ?)")
	checkError(err)
	statement.Exec("Jessica", "McArthur", 30)
	db.Close()
}

func readDB() {
	db, err := sql.Open("sqlite3", "/tmp/personal.db")
	checkError(err)
	var profile Profile
	rows, err := db.Query("select ProfileId, FirstName, LastName, Age from Profile where FirstName = ? and LastName = ?", "Tarik", "Guney")
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&profile.ProfileId, &profile.FirstName, &profile.LastName, &profile.Age)
		checkError(err)
		fmt.Println(profile)
	}
	rows.Close()
	db.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
