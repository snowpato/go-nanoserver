package main

import (
	"fmt"
	"log"
	"net/http"
)

//This application will run a micro server (only HTML) in localhost:8081.
//You can change the port below, or you can parametrize the entry.
//The code is base in the following post:
//https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

var port = "8081"

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../app/" + r.URL.Path[1:])
	})
	
	log.Printf("Running Server in port %s.", port)
	log.Println("To stop the sever, close this app.")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil)) //
}