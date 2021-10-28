package main

import (
	"fmt"
	"net/http"
)

func conexao(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func main() {

	http.HandleFunc("/", conexao)
	http.ListenAndServe(":8080", nil) // indicando a porta em que a nossa aplicação rodará
	fmt.Println("SerVidor Funcionando!")
}
