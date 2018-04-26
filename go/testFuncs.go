package main

import(
 "libs/sqlConnect"
 "libs/howManyFollowers"
 "fmt"
)
func main(){
 sqlConnect.SqlConnect("root","D3m0S14y3c","172.17.2.168","3306","municipios")
 numFollowers := howManyFollowers.GetNumFollowers("tyI3dExsT4Q5LjiA5Oxs9r2rL","t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i","119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw","israel_kastillo")
 fmt.Println(numFollowers)
 fmt.Println(numFollowersTest)
}
