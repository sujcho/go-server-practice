package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//	log.Fatal(http.ListenAndServe(":8080", nil))

	//RequestBulkIndex()

	//client := &http.Client{}
	//RequestMapping(client)
	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			Agg()
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	//RequestAgg(client)

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
