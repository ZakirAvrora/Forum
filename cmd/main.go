package main

import (
	"Forum/internals/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var globalSessions *sessions.Manager

func init() {
	globalSessions = sessions.NewManager("memory", "goSessionId", 3600)
}

func main() {

	//fmt.Println("Connection")
	//
	//db, _ := sql.Open("sqlite3", "./forum.db")
	//statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people(id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	//statement.Exec()
	//
	//statement, _ = db.Prepare("INSERT INTO people(firstname, lastname) VALUES (?, ?)")
	//statement.Exec("Maria", "Roboy")
	//
	//rows, _ := db.Query("SELECT id, firstname, lastname FROM people")
	//var id int
	//var firstname string
	//var lastname string
	//
	//for rows.Next() {
	//	rows.Scan(&id, &firstname, &lastname)
	//	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	//}

}
