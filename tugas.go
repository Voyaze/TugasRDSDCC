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
		fmt.Fprintf(w, "Welcome to the Home Page 2.0")
	})

	http.HandleFunc("/gameinfo", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("gameId")
		getGameInfo(id)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getGameInfo(gameId string) {
	rows, err := Database.Query("SELECT * FROM GameInformation WHERE GameID = ?", gameId)
	if err != nil {
		fmt.Println("Error in Query")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var gameID string
		var gameName string
		var gameGenre string
		err = rows.Scan(&gameID, &gameName, &gameGenre)
		if err != nil {
			fmt.Println("Error in Scan")
			panic(err)
		}
		fmt.Println("Game ID: ", gameID)
		fmt.Println("Game Name: ", gameName)
		fmt.Println("Game Genre: ", gameGenre)
	}
}
