package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {

	http.HandleFunc("/languages", getLanguages)
	http.ListenAndServe(":8080", nil)

}

func getLanguages(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	log.Println("Request Iniciada")
	defer log.Println("Request Finalizada")

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/languages_docker")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	languages, err := selectAllLanguages(db)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(languages)

	ctx.Done()
}

func selectAllLanguages(db *sql.DB) ([]Language, error) {

	rows, err := db.Query("select id, name from languages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []Language
	for rows.Next() {
		var l Language
		err = rows.Scan(&l.ID, &l.Name)
		if err != nil {
			return nil, err
		}

		languages = append(languages, l)
	}

	return languages, nil
}
