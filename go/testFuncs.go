package main

import(
 "libs/sqlConnect"
 "libs/getNumFollowers"
 "libs/getFollowersTest"
 "fmt"
)
func main(){
 sqlConnect.SqlConnect("root","D3m0S14y3c","172.17.2.168","3306","municipios")
 numFollowers := getNumFollowers.HowManyFollowers("tyI3dExsT4Q5LjiA5Oxs9r2rL","t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i","119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw")
 numFollowersTest := getFollowersTest.GetNumFollowers("tyI3dExsT4Q5LjiA5Oxs9r2rL","t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i","119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw","brozoxmiswebs")
 fmt.Println(numFollowers)
 fmt.Println(numFollowersTest)
}
