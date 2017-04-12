package main

import (
	"fmt"
	"net/http"
)

//function callback for url endpoint
func Index(w http.ResponseWriter, r *http.Request) {
	indexed := RequestBulkIndex()

	//fmt.Fprintf(w, "Inputs are sent to Elastic Search, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintln(w, "Inputs are sent to Elastic Search")
	fmt.Fprintln(w, "Indexed documents: ")
	for _, index := range indexed {
		fmt.Fprintf(w, "%s\n", index.Id)
	}
}
func Aggregation(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	var result string
	RequestMapping(client)
	RequestAgg(client, &result)

	fmt.Fprintf(w, "%s", result)

}
