package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/header", helloHandle)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "this is index")
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k,v := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
	}
}

