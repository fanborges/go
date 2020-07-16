package main


import (

    "./fileread"
    "fmt"

)


func main(){
    a, _ := fileread.ListFiles(".")
    fmt.Println(a)
}
