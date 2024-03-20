package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type cliente struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var clientes []cliente

func pegarCliente(w http.ResponseWriter, r *http.Request) {
	//Pegue o parâmetro da solicitação
	params := mux.Vars(r)
	//Percorra todos os clientes e encontre o ID correspondente
	for _, item := range clientes {
		if item.ID == params["id"] {
			//retorne o cliente como JSON
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//se o cliente não for encontrado, retorne um erro 404
	json.NewEncoder(w).Encode(&cliente{})

}

func deletarCliente(w http.ResponseWriter, r *http.Request) {
	//Pegue o parâmetro da solicitação
	params := mux.Vars(r)
	//Percorra todos os clientes e encontre o ID correspondente
	for index, item := range clientes {
		if item.ID == params["id"] {
			//remova o cliente do slice
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(clientes)
}

func criarCliente(w http.ResponseWriter, r *http.Request) {
	// Pegue o parâmetro da solicitação
	params := mux.Vars(r)
	var cliente cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cliente.ID = params["id"]
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(clientes)
}

func main() {
	router := mux.NewRouter()
	//crie alguns clientes
	clientes = append(clientes, cliente{ID: "1", Nome: "João", Email: "teste@email.com"})
	clientes = append(clientes, cliente{ID: "2", Nome: "Maria", Email: "teste2@email.com"})

	router.HandleFunc("/clientes/{id}", pegarCliente).Methods("GET")
	router.HandleFunc("/clientes/{id}", deletarCliente).Methods("DELETE")
	router.HandleFunc("/clientes{id}", criarCliente).Methods("POST")

	http.ListenAndServe(":8000", router)
}
