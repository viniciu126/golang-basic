package main

import (
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Olá mundo!"))
}

func usuarios(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Carregar página de usuários!!"))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
