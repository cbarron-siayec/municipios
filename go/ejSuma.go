package main
// ESTE TIPO DE CODIGO GENERA UN EJECUTABLE, POR LO TANTO DEBE DE DECLARARSE COMO "package main" y tener una funcion "main()"
import(
  "fmt" //FUNCION ESTANDAR DE GO PARA IMPRIMIR
  _ "github.com/dghubble/oauth1" //LIBRERIA EXTERNA Y TIENE GUION BAJO POR QUE NO LO VAMOS A USAR (NO RECOMENDADO)
  //"libs/getAllCandidatesTwitterData" LIBRERIA GRUPO SIAYEC
    "strconv"
)

type PersonaOficina struct{
    nombre  string  `json:"nombre"`
    apellido string `json:"apellido"`
    edad    int     `json:"edad"`
}

saul PersonaOficina= new

func sum(a int, b int) int{
  var c int
  c = a + b
  return c
}

func main(){
//  getAllCandidateData.GetAllCandidateData(546)
    var suma int
    suma = sum(5,6)
    sumaString := strconv.Itoa(suma)
    //fmt.Println(getAllCandidatesTwitterData.GetAllCandidatesTwitterData(213745334))
    fmt.Println(suma)
    
}

package main

import (
	"fmt"
	"strconv"
)

type personaOficina struct{
	nombre	 []string
	apellido []string
	edad	 []int
}


func main() {
	saul := personaOficina{nombre:"Saul", apellido:"Zamora", edad:35}
	saulEdadString := strconv.Itoa(saul.edad)
	fmt.Println("Nombre: "+saul.nombre+" Apellido: "+saul.apellido+" Edad: "+saulEdadString)
}
