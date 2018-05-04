package main
// ESTE TIPO DE CODIGO GENERA UN EJECUTABLE, POR LO TANTO DEBE DE DECLARARSE COMO "package main" y tener una funcion "main()"
import(
  "fmt" //FUNCION ESTANDAR DE GO PARA IMPRIMIR
  _ "github.com/dghubble/oauth1" //LIBRERIA EXTERNA Y TIENE GUION BAJO POR QUE NO LO VAMOS A USAR (NO RECOMENDADO)
  // "libs/getAllCandidateData" LIBRERIA GRUPO SIAYEC
)

func sum(a int, b int) int{
  var c int
  d := a + b
  c = a + b
  return c
}

func main(){
//  getAllCandidateData.GetAllCandidateData(546)
    var suma int
    suma = sum(5,6)
}

