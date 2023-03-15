package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"

	"github.com/captain-corgi/go-service-weaver-example/internal/svc/reverser"
)

func main() {
	// Define configurations
	url := "localhost:12345"
	opts := weaver.ListenerOptions{
		LocalAddress: url,
	}

	// Init root executor
	root := weaver.Init(context.Background())
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("hello listener available on %v\n", lis)

	// Get a client to the Reverser component
	reverserComponent, err := weaver.Get[reverser.Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	// Sample http component
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverserComponent.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})

	// Start http server
	http.Serve(lis, nil)
}
