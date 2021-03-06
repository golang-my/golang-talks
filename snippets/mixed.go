package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	person := mux.Vars(r)["person"]
	w.Write([]byte(fmt.Sprintf("Hello, %s!", person)))
}

func main() {
	// http.Handler
	r := mux.NewRouter()
	r.HandleFunc("/{person}", SayHello)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}
