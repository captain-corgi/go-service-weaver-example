package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	url := "localhost:12345"

	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{
		LocalAddress: url,
	}

	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("hello listener available on %v\n", lis)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
	})
	http.Serve(lis, nil)
}
