package reader

//DATASET:
//https://www.datosabiertos.gob.pe/dataset/inventario-de-recursos-tur%C3%ADsticos

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	NUM_CPU = 4
)

type Recurso struct {
	REGIÓN             string
	PROVINCIA          string
	DISTRITO           string
	Codigo_del_Recurso int
	Nombre_del_Recurso string
	CATEGORIA          string
	Tipo_de_Categoria  string
	Sub_tipo_Categoria string
	URL                string
	LATITUD            *float64
	LONGITUD           *float64
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Comma = ','
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	data, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return data, nil
}
func convertThread(tuples [][]string, ch chan Recurso) {
	var latitud, longitud *float64 = new(float64), new(float64)
	for _, value := range tuples {
		codigo, err := strconv.Atoi(strings.Replace(strings.TrimSpace(value[3]), ",", "", -1))
		check(err)
		if value[9] != "" {
			aux, err := strconv.ParseFloat(strings.TrimSpace(value[9]), 64)
			check(err)
			latitud = &aux
		} else {
			latitud = nil
		}
		if value[10] != "" {
			aux, err := strconv.ParseFloat(strings.TrimSpace(value[10]), 64)
			check(err)
			longitud = &aux
		} else {
			longitud = nil
		}
		recursosingle := Recurso{
			REGIÓN:             value[0],
			PROVINCIA:          value[1],
			DISTRITO:           value[2],
			Codigo_del_Recurso: codigo,
			Nombre_del_Recurso: value[4],
			CATEGORIA:          value[5],
			Tipo_de_Categoria:  value[6],
			Sub_tipo_Categoria: value[7],
			URL:                value[8],
			LATITUD:            latitud,
			LONGITUD:           longitud,
		}
		//put to channel
		ch <- recursosingle
		//recursos = append(recursos, recursosingle)
	}
	close(ch)
}
func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func LoadRecursos() []Recurso {
	data, err := readCSVFromUrl("https://raw.githubusercontent.com/Alextron0102/TA2-Go-API/main/files/Inventario_recursos_turisticos.csv")
	//tuple, err := readData("./files/Inventario_recursos_turisticos.csv")
	check(err)
	channels := make([]chan Recurso, NUM_CPU+1)
	limit := len(data) / NUM_CPU
	fmt.Print("lineas en total: ")
	fmt.Println(len(data))
	iteratoraux := 0
	for i := 0; i < len(data); i += limit {
		chunk := data[i:Min(i+limit, len(data))]
		channels[iteratoraux] = make(chan Recurso)
		go convertThread(chunk, channels[iteratoraux])
		iteratoraux++
	}
	var recursos []Recurso
	for _, channel := range channels {
		for recurso := range channel {
			//PrintRecurso(recurso)
			recursos = append(recursos, recurso)
		}
	}
	return recursos
}

//*--Single threaded load--*
// func LoadRecursos() []Recurso {
// 	tuple, err := readCSVFromUrl("https://raw.githubusercontent.com/Alextron0102/TA2-Go-API/main/files/Inventario_recursos_turisticos.csv")
// 	//tuple, err := readData("./files/Inventario_recursos_turisticos.csv")
// 	check(err)
// 	var latitud, longitud *float64 = new(float64), new(float64)
// 	var recursos []Recurso
// 	for _, value := range tuple {
// 		codigo, err := strconv.Atoi(strings.Replace(strings.TrimSpace(value[3]), ",", "", -1))
// 		check(err)
// 		if value[9] != "" {
// 			aux, err := strconv.ParseFloat(strings.TrimSpace(value[9]), 64)
// 			check(err)
// 			latitud = &aux
// 		} else {
// 			latitud = nil
// 		}
// 		if value[10] != "" {
// 			aux, err := strconv.ParseFloat(strings.TrimSpace(value[10]), 64)
// 			check(err)
// 			longitud = &aux
// 		} else {
// 			longitud = nil
// 		}
// 		recursosingle := Recurso{
// 			REGIÓN:             value[0],
// 			PROVINCIA:          value[1],
// 			DISTRITO:           value[2],
// 			Codigo_del_Recurso: codigo,
// 			Nombre_del_Recurso: value[4],
// 			CATEGORIA:          value[5],
// 			Tipo_de_Categoria:  value[6],
// 			Sub_tipo_Categoria: value[7],
// 			URL:                value[8],
// 			LATITUD:            latitud,
// 			LONGITUD:           longitud,
// 		}
// 		recursos = append(recursos, recursosingle)
// 	}
// 	return recursos
// }

//*-- FOR DEBUGING PURPOSES--*
// func PrintRecurso(rec Recurso) {
// 	if rec.LATITUD != nil && rec.LONGITUD != nil {
// 		fmt.Printf("%s %s %s %d %f %f \n",
// 			rec.REGIÓN,
// 			rec.PROVINCIA,
// 			rec.DISTRITO,
// 			rec.Codigo_del_Recurso,
// 			*rec.LATITUD,
// 			*rec.LONGITUD)
// 	} else {
// 		fmt.Printf("%s %s %s %d \n",
// 			rec.REGIÓN,
// 			rec.PROVINCIA,
// 			rec.DISTRITO,
// 			rec.Codigo_del_Recurso)
// 	}
// }

//*-- FOR DEBUGING PURPOSES--*
// func readlines() {
// 	pwd, err := os.Getwd()
// 	check(err)
// 	file, err := os.Open(pwd + "/files/Inventario_recursos_turisticos.csv")
// 	check(err)
// 	scanner := bufio.NewScanner(file)
// 	result := []string{}
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		result = append(result, line)
// 	}
// 	//restamos uno al resultado final porque ese 1 es el header
// 	fmt.Println(len(result) - 1)
// }

//*-- FOR DEBUGING PURPOSES--*
// func readCSVfromFile(fileName string) ([][]string, error) {
// 	f, err := os.Open(fileName)
// 	if err != nil {
// 		return [][]string{}, err
// 	}
// 	defer f.Close()
// 	r := csv.NewReader(f)
// 	// skip first line
// 	if _, err := r.Read(); err != nil {
// 		return [][]string{}, err
// 	}
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		return [][]string{}, err
// 	}
// 	return records, nil
// }
