package main

import (
	m "TA2-GO-API/reader"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var recursos []m.Recurso

func listarTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	jsonBytes, _ := json.MarshalIndent(recursos, "", " ")
	io.WriteString(res, string(jsonBytes))
}

// func saludo(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	jsonBytes, _ := json.MarshalIndent("Hola", "", " ")
// 	io.WriteString(res, string(jsonBytes))
// }

//endpoints:
//REGIÓN,PROVINCIA,DISTRITO,CATEGORIA,Tipo_de_Categoria,Sub_tipo_Categoria
func handleRequest() {
	http.HandleFunc("/", listarTodo)
	//http.HandleFunc("/saludo", saludo)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func main() {
	recursos = m.LoadRecursos()
	handleRequest()
}
