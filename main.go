package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
	> Write and Add a blog
	> Delete a blog
	> Chnage some part of the glog
	> Add People In There
*/

type blog struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello, This is a go server!")
	fmt.Fprint(w, "Hello, This is a go server!")
}

func addBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "Invalid method, it should be POST!")
		return
	}

	var Blog blog

	err := json.NewDecoder(r.Body).Decode(&Blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ERROR=> ", err)
	}
	fmt.Println("URL: ", r.URL)
	fmt.Println("Body: ", r.Body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog Added Successfully!"))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/add", addBlog)
	fmt.Println("Server started!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
