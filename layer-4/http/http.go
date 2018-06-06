package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello sukablyat")
	w.Write([]byte("!!!"))
}

func main(){
	http.HandleFunc("/", handler)

	fmt.Println("starting server on 8080")
	http.ListenAndServe(":8080", nil)
}
