package main

import (
	"fmt"
	"log"
	"net/http"
)

func server1() {

    fmt.Println("server 1 is running")

	http.HandleFunc("/server1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Server 1"))
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}

func server2() {
    fmt.Println("server 2 is running")

	http.HandleFunc("/server2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Server 2"))
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}

func main() {
	go server1()
	go server2()

	select {}
}
