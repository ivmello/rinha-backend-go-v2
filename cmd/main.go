package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /clientes/{id}/transacoes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})
	mux.HandleFunc("/clientes/{id}/extrato", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})
	http.ListenAndServe("localhost:8080", mux)
}
