package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB = nil

func main() {
	Database, err := sql.Open("mysql", "admin:12345678@tcp(databaserds.cgxx59d0ugdf.us-east-1.rds.amazonaws.com:3306)/Games")
	if err != nil {
		fmt.Println("Database is not found")
		panic(err)
	}
	defer Database.Close()

	err = Database.Ping()
	if err != nil {
		fmt.Println("Database is not connected")
		panic(err)
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Home Page 3.0")
	})

	http.HandleFunc("/gameinfo", func(w http.ResponseWriter, r *http.Request) {
		gameId := r.URL.Query().Get("gameId")
		query := "SELECT * FROM GameInformation WHERE GameID = ?"
		rows, err := Database.Query(query, gameId)
		if err != nil {
			fmt.Println("Error in Query")
			return
		}
		defer rows.Close()

		var GameID string
		var GameName string
		var GameGenre string
		if !rows.Next() {
			fmt.Fprintln(w, "No data with that ID")
			return
		}

		err = rows.Scan(&GameID, &GameName, &GameGenre)
		if err != nil {
			fmt.Fprintln(w, "Error in scanning")
			return
		}
		fmt.Fprintln(w, "GameID: ", GameID)
		fmt.Fprintln(w, "GameName: ", GameName)
		fmt.Fprintln(w, "GameGenre: ", GameGenre)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
