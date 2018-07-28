package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello 51reboot.")
}


func main(){
	http.HandleFunc("/",hello)

	err:=http.ListenAndServe("0.0.0.0:9000",nil)
	if err!=nil{
		fmt.Println("listen fail.")
	}
}