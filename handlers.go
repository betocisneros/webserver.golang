package main

import (
	"fmt"
	"net/http"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Testing..")

}
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint de home...")
}
