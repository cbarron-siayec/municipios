package main

import(
 "libs/sqlConnect"
 "libs/getNumFollowers"
 "fmt"
)
func main(){
  sqlConnect.SqlConnect("root","D3m0S14y3c","172.17.2.168","3306","municipios")
  numFollowers := getNoFollowers.HowManyFollowers("U93DMt4ydkuj7KjtTQOH3A7DT","dhpyVkyl6pxRDIvGsTcvCAQ9uNP2WJ9iFIAD2FlYCDY29nat1c","2253585536-EAQIZqyv9Np8V1fJPcthYlT5TPvpacO0MK209j2","OCgcI2poX3Y9fSVM0i2D0tWP9PepXmSiNJBRUMqAnmZU0")
 fmt.Println(numFollowers)
}
