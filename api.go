package main

import (
	m "TA2-GO-API/reader"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

var recursos []m.Recurso
var regiones, provincias, distritos, categorias, tipo_categorias, sub_tipo_categorias []string

func listarRecursos(res http.ResponseWriter, req *http.Request) {
	region := req.FormValue("region")
	provincia := req.FormValue("provincia")
	distrito := req.FormValue("distrito")
	categoria := req.FormValue("categoria")
	tipo_categoria := req.FormValue("tipo_categoria")
	sub_tipo_categoria := req.FormValue("sub_tipo_categoria")
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	io.WriteString(res, "[")
	for i, recurso := range recursos {
		if !strings.EqualFold(recurso.REGIÓN, region) && region != "" {
			continue
		}
		if !strings.EqualFold(recurso.PROVINCIA, provincia) && provincia != "" {
			continue
		}
		if !strings.EqualFold(recurso.DISTRITO, distrito) && distrito != "" {
			continue
		}
		if !strings.EqualFold(recurso.CATEGORIA, categoria) && categoria != "" {
			continue
		}
		if !strings.EqualFold(recurso.Tipo_de_Categoria, tipo_categoria) && tipo_categoria != "" {
			continue
		}
		if !strings.EqualFold(recurso.Sub_tipo_Categoria, sub_tipo_categoria) && sub_tipo_categoria != "" {
			continue
		}
		jsonBytes, _ := json.MarshalIndent(recurso, "", " ")
		io.WriteString(res, string(jsonBytes))
		if i != len(recursos)-1 {
			io.WriteString(res, ",")
		}
	}
	io.WriteString(res, "]")
}

func listarRegiones(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(regiones, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func listarProvincias(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(provincias, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func listarDistritos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(distritos, "", " ")
	io.WriteString(res, string(jsonBytes))
}
func listarCategorias(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(categorias, "", " ")
	io.WriteString(res, string(jsonBytes))
}
func listarTipoCategorias(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(tipo_categorias, "", " ")
	io.WriteString(res, string(jsonBytes))
}
func listarSubTipoCategorias(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	jsonBytes, _ := json.MarshalIndent(sub_tipo_categorias, "", " ")
	io.WriteString(res, string(jsonBytes))
}

//->Mapa LATITUD,LONGITUD
func handleRequest() {
	http.HandleFunc("/listar", listarRecursos)
	http.HandleFunc("/regiones", listarRegiones)
	http.HandleFunc("/provincias", listarProvincias)
	http.HandleFunc("/distritos", listarDistritos)
	http.HandleFunc("/categorias", listarCategorias)
	http.HandleFunc("/tipo_categorias", listarTipoCategorias)
	http.HandleFunc("/sub_tipo_categorias", listarSubTipoCategorias)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func appendIfMissing(strs []string, str string) []string {
	for _, aux := range strs {
		if aux == str {
			return strs
		}
	}
	return append(strs, str)
}
func getUniqueValues() {
	for _, recurso := range recursos {
		regiones = appendIfMissing(regiones, recurso.REGIÓN)
		provincias = appendIfMissing(provincias, recurso.PROVINCIA)
		distritos = appendIfMissing(distritos, recurso.DISTRITO)
		categorias = appendIfMissing(categorias, recurso.CATEGORIA)
		tipo_categorias = appendIfMissing(tipo_categorias, recurso.Tipo_de_Categoria)
		sub_tipo_categorias = appendIfMissing(sub_tipo_categorias, recurso.Sub_tipo_Categoria)
	}
}
func main() {
	recursos = m.LoadRecursos()
	getUniqueValues()
	handleRequest()
}
