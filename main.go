package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func server1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("server 1 is running")

	mux := http.NewServeMux()

	mux.HandleFunc("/server1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here1")
		w.Write([]byte("This is Server 1"))
	})

	log.Fatal(http.ListenAndServe(":8082", mux))
}

func server2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("server 2 is running")

	http.HandleFunc("/server2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here2")
		w.Write([]byte("This is Server 2"))
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go server1(&wg)
	go server2(&wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
