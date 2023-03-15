package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	root := weaver.Init(context.Background())

	root.New()
}
